package main

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err)
	}

	for _, p := range partitions {
		fmt.Println(p.Topic, "Leader:", p.Leader.ID)
		break
	}
}
