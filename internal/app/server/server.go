package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mishakrpv/musiclib/internal/app/domain/song"
	"github.com/mishakrpv/musiclib/internal/app/infrastructure/data/gorm/postgres"
	"github.com/mishakrpv/musiclib/internal/app/infrastructure/services/clients"

	"go.uber.org/zap"
)

type Server struct {
	port int

	songRepo song.Repository

	musicInfoClient clients.MusicInfoClient
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	zap.L().Debug("Port was read from environment variables", zap.Int("port", port))

	musicInfoServiceBaseUrl := os.Getenv("MUSIC_INFO__URL")

	if musicInfoServiceBaseUrl == "" {
		zap.L().Fatal("MUSIC_INFO__URL required but not provided")
	}
	zap.L().Debug("MusicInfoService URL was read from environment variables", zap.String("music_info_url", musicInfoServiceBaseUrl))
	
	newServer := &Server{
		port: port,

		songRepo: postgres.NewSongRepository(),

		musicInfoClient: clients.NewHttpMusicInfoClient(musicInfoServiceBaseUrl),
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