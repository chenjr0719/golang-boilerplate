package log

import (
	"io"
	"os"
	"time"

	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/rs/zerolog"
)

var (
	Logger = NewLogger()
	Panic  = Logger.Panic
	Fatal  = Logger.Fatal
	Error  = Logger.Error
	Warn   = Logger.Warn
	Info   = Logger.Info
	Debug  = Logger.Debug
	Trace  = Logger.Trace
)

func NewLogger() zerolog.Logger {
	var output io.Writer
	if config.Config.Mode == "release" {
		output = os.Stdout
	} else {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}

	logLevel, err := zerolog.ParseLevel(config.Config.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	logger := zerolog.New(output).With().Timestamp().Caller().Logger().Level(logLevel)

	return logger
}
