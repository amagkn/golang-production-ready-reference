package main

import (
	"context"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		// Если топик не определен здесь, он должен определятся в Message
		Balancer: &kafka.LeastBytes{},
		Logger:   logger.Logger(),
	}

	err := w.WriteMessages(context.Background(),
		// Если тема не будет определена, вернется ошибка
		kafka.Message{
			Topic: "my-topic",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "my-topic",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "my-topic",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}
}
