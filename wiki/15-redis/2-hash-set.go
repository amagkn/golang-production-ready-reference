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

	// Установка значений в хэш
	client.HSet(ctx, "user:1", "name", "Alice", "age", 25)

	// Перезапись значения
	client.HSet(ctx, "user:1", "name", "Bob")

	// Установка значения если оно не существует
	client.HSetNX(ctx, "user:1", "name1", "Alice")

	// Получение значения из хэша
	name, err := client.HGet(ctx, "user:1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", name) // Вывод: name: Alice

	// Проверка существования поля в хэше
	exists, err := client.HExists(ctx, "user:1", "age").Result()
	if err != nil {
		panic(err)
	}

	// Инкремент поля в хэше
	client.HIncrBy(ctx, "user:1", "age", 1)

	fmt.Println("age exists:", exists) // Вывод: age exists: true

	// Получение всех полей хэша
	user, err := client.HGetAll(ctx, "user:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user:", user) // Вывод: user: map[name:Alice age:25]

	// Удаление поля из хэша
	client.HDel(ctx, "user:1", "age")

	// Проверка существования поля в хэше
	exists, err = client.HExists(ctx, "user:1", "age").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("age exists:", exists) // Вывод: age exists: false
}
