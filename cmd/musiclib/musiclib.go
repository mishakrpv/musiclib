package main

import (
	"context"
	"encoding/json"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mishakrpv/musiclib/cmd"
	"github.com/mishakrpv/musiclib/pkg/config"
	"github.com/mishakrpv/musiclib/pkg/logger"
	pserver "github.com/mishakrpv/musiclib/pkg/server"
	"github.com/rs/zerolog/log"
)

//	@title			Musiclib API
//	@version		1.0
//	@description	Effective Mobile test task

//	@contact.email	mishavkrpv@gmail.com

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	// config inits
	config := cmd.NewCmdConfiguration()

	ctx := context.Background()
	if err := runCmd(ctx, &config.Configuration); err != nil {
		stdlog.Println(err)
		os.Exit(1)
	}
}

func runCmd(ctx context.Context, cfg *config.Configuration) error {
	logger.SetupLogger(cfg)

	jsonConf, err := json.Marshal(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Could not marshal static configuration")
		log.Debug().Interface("staticConfiguration", cfg).Msg("Static configuration loaded [struct]")
	} else {
		log.Debug().RawJSON("staticConfiguration", jsonConf).Msg("Static configuration loaded [json]")
	}

	svr, err := setupServer(ctx, cfg)
	if err != nil {
		return err
	}

	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	svr.Start(ctx)
	defer svr.Close()

	svr.Wait()
	log.Info().Msg("Shutting down")
	return nil
}

func setupServer(_ context.Context, cfg *config.Configuration) (*pserver.Server, error) {
	httpServer := &http.Server{
		Addr:         net.JoinHostPort(cfg.ServerConfig.Host, cfg.ServerConfig.Port),
		IdleTimeout:  config.DefaultIdleTimeout,
		ReadTimeout:  config.DefaultReadTimeout,
		WriteTimeout: config.DefaultWriteTimeout,
		Handler:      nil,
	}

	return pserver.NewServer(httpServer), nil
}
