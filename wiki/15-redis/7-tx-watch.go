package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	// Watch — это механизм оптимистической блокировки.
	// Он позволяет отслеживать изменения ключей и отменять транзакцию, если ключи были изменены другим клиентом.
	// Когда вам нужно гарантировать, что ключи не были изменены другим клиентом перед выполнением транзакции.
	err := client.Watch(ctx, func(tx *redis.Tx) error {
		// Получаем текущее значение ключа
		value, err := tx.Get(ctx, "key1").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		// Начало транзакции
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			// Изменяем ключ
			pipe.Set(ctx, "key1", value+"_updated", 0)
			return nil
		})
		return err
	}, "key1")
	if err != nil {
		panic(err)
	}

	fmt.Println("Транзакция выполнена успешно")
}
