package main

import (
	"context"
	"hash/fnv"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:        "my-topic",
		Balancer:     &kafka.Hash{Hasher: fnv.New32a()},
		RequiredAcks: kafka.RequireAll,
		ErrorLogger:  logger.Logger(),
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Three!"),
		},
		kafka.Message{
			Key:   []byte("Key-D"),
			Value: []byte("Four!"),
		},
		kafka.Message{
			Key:   []byte("Key-E"),
			Value: []byte("Five!"),
		},
		kafka.Message{
			Key:   []byte("Key-F"),
			Value: []byte("Six!"),
		},
	)
	if err != nil {
		panic(err)
	}

	if err := w.Close(); err != nil {
		panic(err)
	}
}
