package usecase

import (
	"context"
	"fmt"

	"github.com/amagkn/golang-production-ready-reference/pkg/otel/tracer"
	"github.com/amagkn/golang-production-ready-reference/pkg/transaction"
)

func (u *UseCase) OutboxReadAndProduce(ctx context.Context, limit int) (count int, err error) {
	ctx, span := tracer.Start(ctx, "usecase OutboxReadAndProduce")
	defer span.End()

	// В транзакции
	err = transaction.Wrap(ctx, func(ctx context.Context) error {
		// Читаем сообщения из outbox таблицы БД
		msgs, err := u.postgres.ReadOutboxKafka(ctx, limit)
		if err != nil {
			return fmt.Errorf("u.postgres.ReadOutboxKafka: %w", err)
		}

		count = len(msgs)

		// Пишем в Kafka
		err = u.kafka.Produce(ctx, msgs...)
		if err != nil {
			return fmt.Errorf("u.kafka.Produce: %w", err)
		}

		return nil
	})
	if err != nil {
		return count, fmt.Errorf("transaction.Wrap: %w", err)
	}

	return count, nil
}
