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

	// Установка значения
	client.Set(ctx, "key", "value", time.Hour)

	// Установка значения если оно не существует
	client.SetNX(ctx, "key1", "value1", time.Hour)

	// Установка значения если оно существует
	client.SetXX(ctx, "key2", "value2", time.Hour)

	// Получение значения
	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key:", val) // Вывод: key: value

	// Удаление ключа
	client.Del(ctx, "key")
	client.Del(ctx, "key1")
}
