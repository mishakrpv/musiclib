package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/songs", s.GetAllSongsHandler)
	r.GET("/groups/:group_name/songs/:song_name/lyrics/:verse_number", s.GetSongLyricsHandler)

	r.DELETE("/groups/:group_name/songs/:song_name", s.DeleteSongHandler)

	r.PUT("/groups/:group_name/songs/:song_name", s.UpdateSongHandler)

	r.POST("/songs", s.CreateSongHandler)

	return r
}

func (s *Server) GetAllSongsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Server) GetSongLyricsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Server) DeleteSongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Server) UpdateSongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Server) CreateSongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
