package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

var ctx = context.Background()

// At-Least-Once (минимум один раз - гарантированная доставка)
func atLeastOnce() {
	writer := kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "orders",
		Async:        false,            // Синхронная запись
		RequiredAcks: kafka.RequireAll, // Ждём подтверждения от всех реплик (acks=all) min ISR == 2
	}

	_ = writer

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "my-group",
	})
	msg, _ := reader.FetchMessage(ctx) // Fetch вместо Read
	// check idempotency key -> Redis ttl 5 min. Postgres
	reader.CommitMessages(ctx, msg) // Явное подтверждение
}

// At-Most-Once (максимум один раз - может не доставить)
func atMostOnce() {
	writer := kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "metrics",
		Async:        true,              // Асинхронная запись (нет ожидания ack)
		RequiredAcks: kafka.RequireNone, // acks=0
	}

	_ = writer

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "metrics",
		GroupID: "my-group",
	})

	// Консьюмер читает без подтверждения
	msg, _ := reader.ReadMessage(ctx) // Автоматическое подтверждение
	// reader.CommitMessages() не вызываем!

	_ = msg
}
