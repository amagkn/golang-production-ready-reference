package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:   "my-topic",
		GroupID: "fetch-and-commit-1",
		Logger:  logger.Logger(),
	})
	defer r.Close()

	ctx := context.Background()

	for {
		m, err := r.FetchMessage(ctx) // Читаем
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if err = r.CommitMessages(ctx, m); err != nil { // Коммитим offset
			panic(err)
		}
	}
}
