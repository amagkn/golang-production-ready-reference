package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// default: auto.create.topics.enable='true'
	const topic = "my-topic"
	const partition = 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
		kafka.Message{Value: []byte("four!")},
		kafka.Message{Value: []byte("five!")},
	)
	if err != nil {
		panic(err)
	}
}
