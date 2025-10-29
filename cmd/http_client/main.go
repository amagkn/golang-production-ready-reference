package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
	. "github.com/amagkn/golang-production-ready-reference/pkg/profile_client"
)

var ctx = context.Background()

func main() {
	logger.Init(logger.Config{PrettyConsole: true})

	profile := New(Config{Host: "localhost", Port: "8080"})

	profileIDs := make(chan uuid.UUID, 100)

	go func() {
		for range time.Tick(500 * time.Millisecond) {
			id, err := profile.Create(ctx,
				gofakeit.Name(),
				gofakeit.IntRange(18, 120),
				gofakeit.Email(),
				"+7"+gofakeit.Phone(),
			)
			if err != nil {
				log.Error().Err(err).Msg("profile.Create")

				continue
			}

			profileIDs <- id
		}
	}()

	go func() {
		for id := range profileIDs {
			for range 3 {
				_, err := profile.GetProfile(ctx, id.String())
				if err != nil {
					log.Error().Err(err).Msg("profile.GetProfile")

					continue
				}
			}

			name := gofakeit.Name()

			err := profile.Update(ctx, id.String(), &name, nil, nil, nil)
			if err != nil {
				log.Error().Err(err).Msg("profile.Update")

				continue
			}

			err = profile.Delete(ctx, id.String())
			if err != nil {
				log.Error().Err(err).Msg("profile.Delete")
			}

			// Чтобы иногда вылетала 404 ошибка
			if strings.HasPrefix(id.String(), "0") {
				_, _ = profile.GetProfile(ctx, id.String())
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
}
