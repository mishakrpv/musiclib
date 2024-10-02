package server

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/endpoint/commands/song/create"
	"github.com/mishakrpv/musiclib/internal/endpoint/commands/song/update"
	"github.com/mishakrpv/musiclib/internal/endpoint/query"

	"github.com/gin-gonic/gin"
	pagination "github.com/webstradev/gin-pagination"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		paginator := pagination.New("page", "rowsPerPage", "1", "15", 5, 150)

		api.GET("/songs", paginator, s.SongsHandler)
		api.GET("/songs/:song_id/lyrics", s.LyricsHandler)

		api.DELETE("/songs/:song_id", s.DeleteSongHandler)

		api.PUT("/songs/:song_id", s.UpdateSongHandler)

		api.POST("/songs", s.CreateSongHandler)
	}

	return r
}

func (s *Server) SongsHandler(c *gin.Context) {
	handler := query.NewHandler(s.songRepo)

	filter := &query.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		zap.L().Warn("Something went wrong while binding query", zap.Error(err))
	}

	zap.L().Debug("Query bound",
		zap.String("group", filter.GroupName),
		zap.String("song", filter.SongName),
		zap.String("date", filter.ReleaseDate),
		zap.String("text", filter.Text),
		zap.String("link", filter.Link))

	filter.Link, _ = url.QueryUnescape(filter.Link)
	zap.L().Debug("Decode query link", zap.String("link", filter.Link))

	response, err := handler.Execute(filter)
	if err != nil {
		// TODO: map error to proper status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, &response)
}

func (s *Server) LyricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Server) DeleteSongHandler(c *gin.Context) {

	id := c.Param("song_id")

	zap.L().Debug("Param bound", zap.String("id", id))

	err := s.songRepo.Delete(id)
	if err != nil {
		// TODO: map error to proper status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func (s *Server) UpdateSongHandler(c *gin.Context) {
	songId := c.Param("song_id")

	zap.L().Debug("Param bound", zap.String("id", songId))

	request := &update.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(songId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
	}

	s.songRepo.Update(&song.Song{
		Id:          id,
		GroupName:   request.GroupName,
		SongName:    request.SongName,
		ReleaseDate: request.ReleaseDate,
		Text:        request.Text,
		Link:        request.Link,
	})

	c.Status(http.StatusOK)
}

func (s *Server) CreateSongHandler(c *gin.Context) {
	handler := create.NewHandler(s.songRepo, s.musicInfoClient)

	request := &create.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zap.L().Debug("Body bound",
		zap.String("group", request.Group),
		zap.String("song", request.Song))

	response, err := handler.Execute(request)
	if err != nil {
		// TODO: map error to proper status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, &response)
}
