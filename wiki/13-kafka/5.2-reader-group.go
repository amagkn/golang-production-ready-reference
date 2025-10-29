package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:    "my-topic",
		GroupID:  "my-group-1",
		MaxBytes: 10e6, // 10MB
		Logger:   logger.Logger(),
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
