package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

const (
	// DefaultIdleTimeout before closing an idle connection.
	DefaultIdleTimeout = 180 * time.Second

	// DefaultReadTimeout defines the default maximum duration for reading the entire request, including the body.
	DefaultReadTimeout = 60 * time.Second

	// DefaultWriteTimeout is the maximum duration before timing out writes of the response.
	DefaultWriteTimeout = 60 * time.Second
)

type Configuration struct {
	ServerConfig *ServerConfig

	Log *Log
}

type ServerConfig struct {
	Port string `env:"PORT"`
	Host string `env:"HOST"`
}

func LoadServerConfig() *ServerConfig {
	cfg, err := env.ParseAs[ServerConfig]()
	if err != nil {
		panic(err)
	}

	return &cfg
}

type Log struct {
	Level   string `env:"LOG__LEVEL"`
	Format  string `env:"LOG__FORMAT"`
	NoColor bool   `env:"LOG__NO_COLOR"`

	FilePath   string `env:"LOG__FILEPATH"`
	MaxSize    int    `env:"LOG__MAX_SIZE"`
	MaxAge     int    `env:"LOG__MAX_AGE"`
	MaxBackups int    `env:"LOG__MAX_BACKUPS"`
	Compress   bool   `env:"LOG__COMPRESS"`
}

func LoadLog() *Log {
	cfg, err := env.ParseAs[Log]()
	if err != nil {
		panic(err)
	}

	return &cfg
}
