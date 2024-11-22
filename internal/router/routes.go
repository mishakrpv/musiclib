package router

import (
	"net/http"

	_ "github.com/mishakrpv/musiclib/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	pagination "github.com/webstradev/gin-pagination"
)

func (r *Router) RegisterRoutes() http.Handler {
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.engine.Group("/api/v1")
	{
		paginator := pagination.New("page", "songs", "1", "3", 3, 150)

		api.GET("/songs", paginator, r.SongsHandler)
		api.GET("/songs/:song_id/lyrics", paginator, r.LyricsHandler)

		api.DELETE("/songs/:song_id", r.DeleteSongHandler)

		api.PUT("/songs/:song_id", r.UpdateSongHandler)

		api.POST("/songs", r.CreateSongHandler)
	}

	return r.engine
}
