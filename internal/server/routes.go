package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/endpoint/command"
	"github.com/mishakrpv/musiclib/internal/endpoint/query"
	"github.com/mishakrpv/musiclib/internal/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pagination "github.com/webstradev/gin-pagination"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		paginator := pagination.New("page", "songs", "1", "3", 3, 150)

		api.GET("/songs", paginator, s.SongsHandler)
		api.GET("/songs/:song_id/lyrics", paginator, s.LyricsHandler)

		api.DELETE("/songs/:song_id", s.DeleteSongHandler)

		api.PUT("/songs/:song_id", s.UpdateSongHandler)

		api.POST("/songs", s.CreateSongHandler)
	}

	return r
}

func (s *Server) SongsHandler(c *gin.Context) {
	qr := query.NewSongsQuery(s.songRepo)

	filter := &query.Filter{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		zap.L().Warn("Something went wrong while binding query", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	page := c.GetInt("page")
	amount := c.GetInt("songs")
	amount--
	page--

	zap.L().Debug("Query bound",
		zap.String("group", filter.GroupName),
		zap.String("song", filter.SongName),
		zap.String("date", filter.ReleaseDate),
		zap.String("text", filter.Text),
		zap.String("link", filter.Link),
		zap.Int("page", page))

	filter.Link, _ = url.QueryUnescape(filter.Link)
	zap.L().Debug("Decode query link", zap.String("link", filter.Link))

	response, err := qr.Execute(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response = response[page*amount : (page+1)*amount]

	res := make([]interface{}, 0, len(response))
	for _, item := range response {
		if item != nil {
			res = append(res, item)
		}
	}

	c.JSON(http.StatusOK, &res)
}

func (s *Server) LyricsHandler(c *gin.Context) {
	query := query.NewLyricsQuery(s.songRepo)

	id := c.Param("song_id")
	page := c.GetInt("page")
	zap.L().Debug("Params bound", zap.String("id", id), zap.Int("page", page))

	verse, err := query.Execute(id, page)
	if err != nil {
		var status int = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrVerseNotFound) || errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{fmt.Sprintf("verse number %d:", page): *verse})
}

func (s *Server) DeleteSongHandler(c *gin.Context) {
	id := c.Param("song_id")

	zap.L().Debug("Param bound", zap.String("id", id))

	err := s.songRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) UpdateSongHandler(c *gin.Context) {
	songId := c.Param("song_id")

	zap.L().Debug("Param bound", zap.String("id", songId))

	request := &command.UpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(songId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect song id"})
		return
	}

	err = s.songRepo.Update(&song.Song{
		Id:          id,
		GroupName:   request.GroupName,
		SongName:    request.SongName,
		ReleaseDate: request.ReleaseDate,
		Text:        request.Text,
		Link:        request.Link,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) CreateSongHandler(c *gin.Context) {
	cmd := command.NewCreateCommand(s.songRepo, s.musicInfoClient)

	request := &command.CreateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zap.L().Debug("Body bound",
		zap.String("group", request.Group),
		zap.String("song", request.Song))

	response, err := cmd.Execute(request)
	if err != nil {
		var status = http.StatusInternalServerError
		if errors.Is(err, apperror.ErrSongNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &response)
}
