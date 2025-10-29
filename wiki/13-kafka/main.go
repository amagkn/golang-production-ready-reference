package main

import (
	"context"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	w := &kafka.Writer{ // Producer
		Addr:         kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:        "my-topic",
		Logger:       logger.Logger(),
		RequiredAcks: kafka.RequireAll,
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("One!"),
		},
	)
	if err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}
}
