package logging

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Configuration for logging
type Config struct {
	ConsoleLoggingEnabled bool   // Enable console logging
	EncodeLogsAsJson      bool   // EncodeLogsAsJson makes the log framework log JSON
	FileLoggingEnabled    bool   // FileLoggingEnabled makes the framework log to a file the fields below can be skipped if this value is false!
	Directory             string // Directory to log to to when filelogging is enabled
	Filename              string // Filename is the name of the logfile which will be placed inside the directory
	MaxSize               int    // MaxSize the max size in MB of the logfile before it's rolled
	MaxBackups            int    // MaxBackups the max number of rolled files to keep
	MaxAge                int    // MaxAge the max age in days to keep a logfile
}

var zlog *zerolog.Logger

// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func (cfg *Config) Init() *zerolog.Logger {
	var writers []io.Writer
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	setupLevel(os.Getenv("LOG_LEVEL"))

	if cfg.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if cfg.FileLoggingEnabled {
		writers = append(writers,
			&lumberjack.Logger{
				Filename:   path.Join(cfg.Directory, cfg.Filename),
				MaxBackups: cfg.MaxBackups, // files
				MaxSize:    cfg.MaxSize,    // megabytes
				MaxAge:     cfg.MaxAge,     // days
			})
	}
	mw := io.MultiWriter(writers...)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()
	logger.Info().
		Bool("fileLogging", cfg.FileLoggingEnabled).
		Bool("jsonLogOutput", cfg.EncodeLogsAsJson).
		Str("logDirectory", cfg.Directory).
		Str("fileName", cfg.Filename).
		Int("maxSizeMB", cfg.MaxSize).
		Int("maxBackups", cfg.MaxBackups).
		Int("maxAgeInDays", cfg.MaxAge).
		Msg("logging configured")
	zlog = &logger
	return zlog
}

func setupLevel(level string) {
	switch level {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
}

func Get() *zerolog.Logger {
	return zlog
}
