package config

import "time"

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
	Port string
	Host string
}

type Log struct {
	Level   string
	Format  string
	NoColor bool

	FilePath   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}
