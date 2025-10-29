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

	// Push:
	// Добавление элементов в конец списка
	client.RPush(ctx, "mylist", "one", "two", "three")

	// Добавление элементов в начало списка
	client.LPush(ctx, "mylist", "zero")

	// Получение всех элементов списка
	list, err := client.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", list) // Вывод: list: [one two three]

	// Получение элемента по индексу
	item, err := client.LIndex(ctx, "mylist", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("item:", item) // Вывод: item: two

	// Удаление элемента по значению
	client.LRem(ctx, "mylist", 1, "two")

	// Pop:
	// Удаление и получение первого элемента
	item, err = client.RPop(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("item:", item) // Вывод: item: zero

	// Удаление и получение последнего элемента
	item, err = client.RPop(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("item:", item) // Вывод: item: three

	// RPushX:
	// Добавление элемента в конец списка, если список существует
	client.RPushX(ctx, "mylist", "four")

	// Получение количества элементов в списке
	length, err := client.LLen(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length:", length)

	// Удаление списка
	client.Del(ctx, "mylist")
}
