package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/amagkn/golang-production-ready-reference/internal/adapter/kafka_producer"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/grpc"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/kafka_consumer"
	"github.com/amagkn/golang-production-ready-reference/internal/controller/worker"
	"github.com/amagkn/golang-production-ready-reference/pkg/httpserver"
	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
	"github.com/amagkn/golang-production-ready-reference/pkg/otel"
	"github.com/amagkn/golang-production-ready-reference/pkg/postgres"
	"github.com/amagkn/golang-production-ready-reference/pkg/redis"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App           App
	HTTP          httpserver.Config
	GRPC          grpc.Config
	Logger        logger.Config
	OTEL          otel.Config
	Postgres      postgres.Config
	Redis         redis.Config
	KafkaConsumer kafka_consumer.Config
	KafkaProducer kafka_producer.Config
	OutboxKafka   worker.OutboxKafkaConfig
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
