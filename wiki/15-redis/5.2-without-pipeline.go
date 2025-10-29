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

	now := time.Now()

	// Без Pipeline
	// Клиент отправляет 100_000 запросов по одному и ждет ответа на каждый.
	for i := 0; i < 100_000; i++ {
		client.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), time.Minute)
	}

	fmt.Println(time.Since(now))
}
