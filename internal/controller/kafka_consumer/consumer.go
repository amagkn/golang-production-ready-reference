package kafka_consumer

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"

	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"

	"github.com/amagkn/golang-production-ready-reference/internal/usecase"
	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
	"github.com/amagkn/golang-production-ready-reference/pkg/otel/tracer"
)

type Config struct {
	Addr  []string `envconfig:"KAFKA_CONSUMER_ADDR" required:"true"`
	Topic string   `default:"awesome-topic"         envconfig:"KAFKA_CONSUMER_TOPIC"`
	Group string   `default:"awesome-group"         envconfig:"KAFKA_CONSUMER_GROUP"`
}

type Consumer struct {
	config  Config
	reader  *kafka.Reader
	usecase *usecase.UseCase
	stop    context.CancelFunc
	done    chan struct{}
}

func New(cfg Config, uc *usecase.UseCase) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        cfg.Addr,
		Topic:          cfg.Topic,
		GroupID:        cfg.Group,
		ErrorLogger:    logger.ErrorLogger(),
		CommitInterval: 100 * time.Millisecond,
	})

	ctx, stop := context.WithCancel(context.Background())

	c := &Consumer{
		config:  cfg,
		reader:  r,
		usecase: uc,
		stop:    stop,
		done:    make(chan struct{}),
	}

	go c.run(ctx)

	return c
}

func (c *Consumer) run(ctx context.Context) {
	log.Info().Msg("kafka consumer: started")

FOR:
	for {
		// Читаем сообщение из Kafka
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			switch {
			case errors.Is(err, context.Canceled):
				log.Info().Msg("kafka consumer: context canceled")

				break FOR
			case errors.Is(err, io.EOF):
				log.Warn().Err(err).Msg("kafka consumer: FetchMessage")

				break FOR
			}

			log.Error().Err(err).Msg("kafka consumer: FetchMessage")
		}

		ctx, span := tracer.Start(ctx, "kafka consumer from "+c.config.Topic,
			trace.WithSpanKind(trace.SpanKindConsumer),
			trace.WithAttributes(
				semconv.MessagingSystemKafka,
				semconv.MessagingDestinationSubscriptionName(c.config.Topic),
				semconv.MessagingConsumerGroupName(c.config.Group),
				semconv.MessagingKafkaMessageKey(string(m.Key)),
			),
		)

		log.Info().Str("key", string(m.Key)).Msg("kafka consumer: message received")

		// Тут вызываем метод из usecase для обработки сообщения

		// Коммитим оффсет в consumer group
		if err = c.reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka consumer: CommitMessages")
		}

		span.End() // Закрываем span
	}

	close(c.done)
}

func (c *Consumer) Close() {
	log.Info().Msg("kafka consumer: closing")

	c.stop()

	if err := c.reader.Close(); err != nil {
		log.Error().Err(err).Msg("kafka consumer: reader.Close")
	}

	<-c.done

	log.Info().Msg("kafka consumer: closed")
}
