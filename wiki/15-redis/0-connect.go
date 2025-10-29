package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Создаем клиент Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес Redis
		Password: "",               // Пароль (если есть)
		DB:       0,                // Номер базы данных
	})

	ctx := context.Background()

	// Проверяем подключение
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong) // Вывод: PONG
}
