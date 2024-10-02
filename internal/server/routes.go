package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mishakrpv/musiclib/internal/endpoint/commands/song/create"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/songs", s.SongsHandler)
		api.GET("/groups/:group_name/songs/:song_name/lyrics/:verse_number", s.LyricsHandler)

		api.DELETE("/groups/:group_name/songs/:song_name", s.DeleteSongHandler)

		api.PUT("/groups/:group_name/songs/:song_name", s.UpdateSongHandler)

		api.POST("/songs", s.CreateSongHandler)
	}

	return r
}

func (s *Server) SongsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
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

	zap.L().Debug("Body binded",
		zap.String("group", request.Group),
		zap.String("song", request.Song))

	response, err := handler.Execute(request)
	if err != nil {
		// TODO: map error to proper status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, &response)
}
