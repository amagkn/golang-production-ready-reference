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

	// Создание pipeline
	pipe := client.Pipeline()

	// Клиент отправляет все 100_000 команд за один раз, а затем получает все ответы
	for i := 0; i < 100_000; i++ {
		pipe.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), time.Minute)
	}

	// Выполнение команд
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(now))
}
