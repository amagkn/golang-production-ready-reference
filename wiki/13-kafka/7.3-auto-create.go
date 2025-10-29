package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"kafka/pkg/logger"
)

func main() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:                  "topic-A",
		AllowAutoTopicCreation: true,
		Logger:                 logger.Logger(),
	}

	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	}

	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
		break
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
