package server

import (
	"net/http"
	"net/url"

	"github.com/mishakrpv/musiclib/internal/endpoint/commands/song/create"
	"github.com/mishakrpv/musiclib/internal/endpoint/query"

	"github.com/gin-gonic/gin"
	"github.com/webstradev/gin-pagination"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		paginator := pagination.New("page", "rowsPerPage", "1", "15", 5, 150)
		
		api.GET("/songs", paginator, s.SongsHandler)
		api.GET("/groups/:group_name/songs/:song_name/lyrics/:verse_number", s.LyricsHandler)

		api.DELETE("/groups/:group_name/songs/:song_name", s.DeleteSongHandler)

		api.PUT("/groups/:group_name/songs/:song_name", s.UpdateSongHandler)

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
	c.JSON(http.StatusOK, "")
}

func (s *Server) UpdateSongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
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
