package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Без использования Pipeline:
//     - Клиент отправляет команду на сервер.
//     - Ждет ответа от сервера.
//     - Отправляет следующую команду.
// С использованием Pipeline:
//     - Клиент отправляет несколько команд сразу.
//     - Сервер обрабатывает их и возвращает все ответы за один раз.

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// Создание pipeline
	pipe := client.Pipeline()

	// Добавление команд
	pipe.Set(ctx, "key1", "value1", time.Minute)
	pipe.Set(ctx, "key2", "value2", time.Minute)

	// Выполнение команд
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pipeline executed")
}
