package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/amagkn/golang-production-ready-reference/config"
	"github.com/amagkn/golang-production-ready-reference/internal/adapter/kafka_producer"
	"github.com/amagkn/golang-production-ready-reference/internal/adapter/postgres"
	"github.com/amagkn/golang-production-ready-reference/internal/adapter/redis"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/grpc"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/http"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/kafka_consumer"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/worker"
	"github.com/amagkn/golang-production-ready-reference/internal/usecase"
	"github.com/amagkn/golang-production-ready-reference/pkg/httpserver"
	"github.com/amagkn/golang-production-ready-reference/pkg/metrics"
	pgpool "github.com/amagkn/golang-production-ready-reference/pkg/postgres"
	redislib "github.com/amagkn/golang-production-ready-reference/pkg/redis"
	"github.com/amagkn/golang-production-ready-reference/pkg/router"
	"github.com/amagkn/golang-production-ready-reference/pkg/transaction"
)

func Run(ctx context.Context, c config.Config) error {
	// Postgres
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	transaction.Init(pgPool)

	// Redis
	redisClient, err := redislib.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redislib.New: %w", err)
	}

	// Kafka producer
	kafkaProducer := kafka_producer.New(c.KafkaProducer, metrics.NewProcess())

	// UseCase
	uc := usecase.New(
		postgres.New(),
		redis.New(redisClient),
		kafkaProducer,
	)

	// Kafka consumer
	kafkaConsumer := kafka_consumer.New(c.KafkaConsumer, uc)

	// Outbox Kafka worker
	outboxKafkaWorker := worker.NewOutboxKafka(uc, c.OutboxKafka)

	// Metrics
	httpMetrics := metrics.NewHTTPServer()

	// GRPC
	grpcServer, err := grpc.New(c.GRPC, uc)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	// HTTP
	r := router.New()
	http.ProfileRouter(r, uc, httpMetrics)
	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig // wait signal

	log.Info().Msg("app: got signal to stop")

	// Controllers close
	httpServer.Close()
	grpcServer.Close()
	outboxKafkaWorker.Close()
	kafkaConsumer.Close()

	// Adapters close
	redisClient.Close()
	kafkaProducer.Close()
	pgPool.Close()

	log.Info().Msg("app: stopped")

	return nil
}
