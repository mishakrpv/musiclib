package main

import (
	"github.com/mishakrpv/musiclib/internal/pkg/app"
	
	"go.uber.org/zap"
)

func main() {
	app := app.NewApp()

	err := app.Run()
	if err != nil {
		zap.L().Fatal("An error occured while starting application", zap.Error(err))
	}
}
