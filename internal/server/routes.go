package server

import (
	"net/http"

	_ "github.com/mishakrpv/musiclib/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pagination "github.com/webstradev/gin-pagination"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
