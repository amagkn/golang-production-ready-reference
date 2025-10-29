package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:   "my-topic",
		GroupID: "auto-commit-1",
		Logger:  logger.Logger(),
	})
	defer r.Close()

	now := time.Now()

	for {
		fmt.Println(time.Since(now))

		m, err := r.ReadMessage(context.Background()) // Читаем и сразу коммитим offset
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
