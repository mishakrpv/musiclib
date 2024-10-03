package main

import (
	"github.com/mishakrpv/musiclib/internal/server"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

//	@title			Musiclib API
//	@version		1.0
//	@description	Effective Mobile test task

//	@contact.email	mishavkrpv@gmail.com

//	@host		localhost:8080
//	@BasePath	/api/v1
func main() {
	server.ConfigureLogging()

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
