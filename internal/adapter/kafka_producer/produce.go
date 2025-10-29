package kafka_producer

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"github.com/amagkn/golang-production-ready-reference/internal/domain"
	"github.com/amagkn/golang-production-ready-reference/pkg/logger"
	"github.com/amagkn/golang-production-ready-reference/pkg/metrics"
	"github.com/amagkn/golang-production-ready-reference/pkg/otel/tracer"
)

type Config struct {
	Addr []string `envconfig:"KAFKA_WRITER_ADDR" required:"true"`
}

type Producer struct {
	config  Config
	writer  *kafka.Writer
	metrics *metrics.Process
}

func New(c Config, m *metrics.Process) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(c.Addr...),
		RequiredAcks: kafka.RequireAll,
		ErrorLogger:  logger.ErrorLogger(),
		Async:        true,
	}

	return &Producer{
		config:  c,
		writer:  w,
		metrics: m,
	}
}

func (p *Producer) Produce(ctx context.Context, events ...domain.Event) error {
	ctx, span := tracer.Start(ctx, "adapter kafka Produce")
	defer span.End()

	const produce = "produce"

	defer p.metrics.Duration(produce, time.Now())

	var msgs []kafka.Message

	for _, e := range events {
		msg := kafka.Message{
			Topic: e.Topic,
			Key:   e.Key,
			Value: e.Value,
		}

		msgs = append(msgs, msg)
	}

	err := p.writer.WriteMessages(ctx, msgs...)
	if err != nil {
		p.metrics.Total(produce, metrics.Error)

		return fmt.Errorf("p.writer.WriteMessages: %w", err)
	}

	p.metrics.Total(produce, metrics.Ok)

	return nil
}

func (p *Producer) Close() {
	err := p.writer.Close()
	if err != nil {
		log.Error().Err(err).Msg("kafka producer: p.writer.Close")
	}

	log.Info().Msg("kafka producer: closed")
}
