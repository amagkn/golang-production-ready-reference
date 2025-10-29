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

	// Установка ключа с TTL
	err := client.Set(ctx, "tempkey", "tempvalue", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Получение TTL
	ttl, err := client.TTL(ctx, "tempkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("TTL:", ttl) // Вывод: TTL: 10s
}
