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
		Brokers:     []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:       "my-topic",
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		MaxWait:     10 * time.Second,
		Logger:      logger.Logger(),
		ErrorLogger: logger.Logger(),
	})
	defer r.Close()

	//r.SetOffset(2) // сделать сдвиг по оффсету
	r.SetOffsetAt(context.Background(), time.Now().Add(-4*24*time.Hour)) // сделать сдвиг по времени

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		fmt.Printf(m.Time.String())
	}
}
