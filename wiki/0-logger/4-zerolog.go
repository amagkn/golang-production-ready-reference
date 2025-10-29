package main

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var ErrTest = errors.New("test error")

	zerologInit(true, "debug")

	log.Debug().Msg("message")
	log.Info().Msg("message")
	log.Warn().Msg("message")
	log.Error().Err(ErrTest).Msg("message")

	log.Info().
		Str("url", "http://example.com").
		Int("attempt", 3).
		Dur("backoff", time.Second).
		Msg("message")

	log.Fatal().Err(ErrTest).Msg("message")
}

func zerologInit(pretty bool, level string) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		zerolog.SetGlobalLevel(lvl)
	}

	log.Logger = log.With().
		Caller().
		Logger()

	if pretty == true {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("Logger initialized")
}
