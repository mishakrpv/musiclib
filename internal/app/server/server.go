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

var (
	musicInfoServiceBaseUrl = os.Getenv("MUSIC_INFO__URL")
)

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	zap.L().Debug("Port was read from environment variables", zap.Int("port", port))

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
