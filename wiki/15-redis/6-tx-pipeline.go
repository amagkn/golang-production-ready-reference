package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	// TxPipeline — это способ выполнения нескольких команд в Redis атомарно (как одна операция).
	// Все команды, добавленные в TxPipeline, выполняются последовательно при вызове Exec.
	tx := client.TxPipeline()

	// Добавление команд в транзакцию
	tx.Set(ctx, "key1", "value1", 0) // Установка значения
	tx.Set(ctx, "key2", "value2", 0) // Установка значения

	// Операции с множествами
	tx.SAdd(ctx, "myset", "one", "two", "three") // Добавление элементов в множество
	tx.SRem(ctx, "myset", "two")                 // Удаление элемента из множества

	// Операции с хэшами
	tx.HSet(ctx, "myhash", "field1", "value1") // Установка поля в хэше
	tx.HSet(ctx, "myhash", "field2", "value2") // Установка поля в хэше

	// Выполнение транзакции
	_, err := tx.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction executed")
}
