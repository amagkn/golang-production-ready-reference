package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// default: auto.create.topics.enable='true'
	const topic = "my-topic"
	const partition = 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	batch := conn.ReadBatchWith(kafka.ReadBatchConfig{
		MinBytes: 10e3, // 10KB min
		MaxBytes: 1e6,  // 1MB max
		MaxWait:  5 * time.Second,
	})

	b := make([]byte, 10e6) // 10MB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	batch.Close()
}
