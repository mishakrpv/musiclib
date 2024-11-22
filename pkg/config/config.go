package config

type Configuration struct {
	Log *Log
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
