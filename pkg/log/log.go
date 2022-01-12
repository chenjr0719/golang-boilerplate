package log

import (
	"io"
	"os"
	"time"

	"github.com/chenjr0719/golang-boilerplate/pkg/config"
	"github.com/rs/zerolog"
)

func NewLogger(conf config.Config) zerolog.Logger {
	var output io.Writer

	if conf.Mode == "release" {
		output = os.Stdout
	} else {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	}

	logLevel, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	logger := zerolog.New(output).With().Timestamp().Caller().Logger().Level(logLevel)

	return logger
}
