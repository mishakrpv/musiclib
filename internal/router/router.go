package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/pkg/infra/musicinfo"
	"github.com/mishakrpv/musiclib/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Provides routing and handlers.
type Router struct {
	engine *gin.Engine

	musicinfoClient musicinfo.Client
	songRepo        song.Repository
}

// Creates new router.
func New(client musicinfo.Client, songRepo song.Repository) http.Handler {
	engine := gin.New()

	engine.Use(gin.LoggerWithWriter(logger.NoLevel(log.Logger, zerolog.InfoLevel)))

	router := &Router{
		engine:          engine,
		musicinfoClient: client,
		songRepo:        songRepo,
	}

	return router.registerRoutes()
}
