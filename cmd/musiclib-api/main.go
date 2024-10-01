package main

import (
	"os"

	"github.com/mishakrpv/musiclib/internal/app"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	configureLogging()

	server := app.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}

func configureLogging() {
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
