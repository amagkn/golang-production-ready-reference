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

	// Добавление элементов в множество
	client.SAdd(ctx, "myset", "zero", "one", "two", "three")

	// Проверка наличия элемента в множестве
	exists, err := client.SIsMember(ctx, "myset", "two").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("exists:", exists) // Вывод: exists: true

	// Удаляем элемент из множества
	client.SRem(ctx, "myset", "zero")

	// Получение всех элементов множества
	set, err := client.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set:", set) // Вывод: set: [one three two]

	// Создаем второе множество
	client.SAdd(ctx, "myset2", "two", "three", "four")

	// Объединяем два множества
	union, err := client.SUnion(ctx, "myset", "myset2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Объединение множеств:", union) // Вывод: [one three two four]

	// Пересечение множеств
	intersection, err := client.SInter(ctx, "myset", "myset2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Пересечение множеств:", intersection) // Вывод: [three two]

	// Разность множеств
	difference, err := client.SDiff(ctx, "myset", "myset2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Разность множеств:", difference) // Вывод: [one]
}
