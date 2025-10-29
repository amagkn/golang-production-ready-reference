package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// Подписка на канал
	pubsub := client.Subscribe(ctx, "mychannel")
	defer pubsub.Close()

	// Горутина для чтения сообщений
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Println("Received:", msg.Payload)
		}
	}()

	// Публикация сообщения в канал
	client.Publish(ctx, "mychannel", "hello")
	client.Publish(ctx, "mychannel", "world")

	// Ожидание завершения
	time.Sleep(1 * time.Second)
}
