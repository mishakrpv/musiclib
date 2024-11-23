package cmd

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mishakrpv/musiclib/pkg/config"
)

type CmdConfiguration struct {
	config.Configuration
}

func NewCmdConfiguration() *CmdConfiguration {
	return &CmdConfiguration{
		Configuration: config.Configuration{
			ServerConfig: &config.ServerConfig{
				Port: getenvOrDefault("PORT", "8080"),
				Host: getenvOrDefault("HOST", "localhost"),
			},
			Log:          config.Load[config.Log](),
			DBConfig:     config.Load[config.DBConfig](),
			MusicInfoURL: getenvOrDefault("MUSIC_INFO_URL", "http://localhost:3256"),
		},
	}
}

func getenvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return defaultValue
}
