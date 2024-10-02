package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mishakrpv/musiclib/internal/domain/song"
	"github.com/mishakrpv/musiclib/internal/infrastructure/data/gorm"
	"github.com/mishakrpv/musiclib/internal/infrastructure/services"

	"go.uber.org/zap"
)

type Server struct {
	port int

	songRepo song.Repository

	musicInfoClient services.MusicInfoClient
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	zap.L().Debug("Port was read from environment variables", zap.Int("port", port))

	musicInfoBaseUrl := os.Getenv("MUSIC_INFO__URL")

	if musicInfoBaseUrl == "" {
		zap.L().Fatal("MUSIC_INFO__URL required but not provided")
	}
	zap.L().Debug("MusicInfo URL was read from environment variables", zap.String("musicinfo_url", musicInfoBaseUrl))
	
	newServer := &Server{
		port: port,

		songRepo: gorm.NewSongRepository(),

		musicInfoClient: services.NewHttpMusicInfoClient(musicInfoBaseUrl),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func ConfigureLogging() {
	var logger *zap.Logger

	env := os.Getenv("ENV")

	if env == "production" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)

	defer logger.Sync()

	zap.L().Info("Logging has been configured", zap.String("env", env))
}