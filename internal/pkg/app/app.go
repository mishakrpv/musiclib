package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	app := &App{}

	configureLogging()

	app.router = gin.Default()

	return app
}

func (app *App) Run() error {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return app.router.Run(":" + port)
}

func configureLogging() {
	var logger *zap.Logger

	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
		os.Setenv("ENV", env)
	}

	if env == "production" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)

	defer logger.Sync()
}
