package logger

import (
	"io"
	stdlog "log"
	"os"
	"strings"
	"time"

	"github.com/mishakrpv/musiclib/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// hide the first logs before the setup of the logger.
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func SetupLogger(cfg *config.Configuration) {
	// configure log format
	w := getLogWriter(cfg)

	// configure log level
	logLevel := getLogLevel(cfg)

	// create logger
	logCtx := zerolog.New(w).With().Timestamp()
	if logLevel <= zerolog.DebugLevel {
		logCtx = logCtx.Caller()
	}

	log.Logger = logCtx.Logger().Level(logLevel)
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.SetGlobalLevel(logLevel)

	// configure default standard log.
	stdlog.SetFlags(stdlog.Lshortfile | stdlog.LstdFlags)
	stdlog.SetOutput(NoLevel(log.Logger, zerolog.DebugLevel))
}

func getLogWriter(cfg *config.Configuration) io.Writer {
	var w io.Writer = os.Stdout

	if cfg.Log != nil && len(cfg.Log.FilePath) > 0 {
		_, _ = os.OpenFile(cfg.Log.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		w = &lumberjack.Logger{
			Filename:   cfg.Log.FilePath,
			MaxSize:    cfg.Log.MaxSize,
			MaxBackups: cfg.Log.MaxBackups,
			MaxAge:     cfg.Log.MaxAge,
			Compress:   true,
		}
	}

	if cfg.Log == nil || cfg.Log.Format != "json" {
		w = zerolog.ConsoleWriter{
			Out:        w,
			TimeFormat: time.RFC3339,
			NoColor:    cfg.Log != nil && (cfg.Log.NoColor || len(cfg.Log.FilePath) > 0),
		}
	}

	return w
}

func getLogLevel(cfg *config.Configuration) zerolog.Level {
	levelStr := "info"
	if cfg.Log != nil && cfg.Log.Level != "" {
		levelStr = strings.ToLower(cfg.Log.Level)
	}

	logLevel, err := zerolog.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		log.Error().Err(err).
			Str("logLevel", levelStr).
			Msg("Unspecified or invalid log level, setting the level to default (ERROR)...")

		logLevel = zerolog.ErrorLevel
	}

	return logLevel
}
