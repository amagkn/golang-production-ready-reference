package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger() KafkaLogger {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})

	return KafkaLogger{logger: log.Logger}
}

type KafkaLogger struct {
	logger zerolog.Logger
}

func (l KafkaLogger) Printf(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}
