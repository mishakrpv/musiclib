package main

import (
	"github.com/mishakrpv/musiclib/internal/server"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	server.ConfigureLogging()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
