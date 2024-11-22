package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/pkg/infra/musicinfo"
)

type Router struct {
	engine *gin.Engine

	musicinfoClient musicinfo.Client
	songRepo        song.Repository
}

func New(client musicinfo.Client, songRepo song.Repository) http.Handler {
	router := &Router{
		engine:          gin.New(),
		musicinfoClient: client,
		songRepo:        songRepo,
	}

	return router.RegisterRoutes()
}
