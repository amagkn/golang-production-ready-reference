package main

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type ctxKey struct{}

func main() {
	ctx := context.Background()

	logger := zap.NewExample()

	logger = logger.With(zap.String("version", "v0.1.0"))

	for id := range 5 {
		go Controller(ctx, logger, id)
	}

	time.Sleep(time.Second)
}

func Controller(ctx context.Context, logger *zap.Logger, requestID int) {
	logger = logger.With(zap.Int("requestID", requestID))

	ctx = context.WithValue(ctx, ctxKey{}, logger)

	Handler(ctx)
}

func Handler(ctx context.Context) {
	UseCase(ctx)
}

func UseCase(ctx context.Context) {
	Adapter(ctx)
}

func Adapter(ctx context.Context) {
	logger, ok := ctx.Value(ctxKey{}).(*zap.Logger)
	if ok {
		logger.Info("Adapter call!")
	}
}
