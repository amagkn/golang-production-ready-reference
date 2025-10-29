package kafka_consumer

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"github.com/amagkn/golang-production-ready-reference/internal/usecase"
	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
)

type Config struct {
	Addr  []string `envconfig:"KAFKA_CONSUMER_ADDR"   required:"true"`
	Topic string   `envconfig:"KAFKA_CONSUMER_TOPIC"  default:"awesome-topic"`
	Group string   `envconfig:"KAFKA_CONSUMER_GROUP"  default:"awesome-group"`
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

		log.Info().Str("key", string(m.Key)).Msg("kafka consumer: message received")

		// Тут вызываем метод из usecase для обработки сообщения

		// Коммитим оффсет в consumer group
		if err = c.reader.CommitMessages(ctx, m); err != nil {
			log.Error().Err(err).Msg("kafka consumer: CommitMessages")
		}
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
