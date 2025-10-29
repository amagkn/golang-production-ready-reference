package main

import (
	"context"
	"fmt"
)

type contextKey struct{}
type notFoundKey struct{}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, contextKey{}, "SUCCESS!")

	v := ctx.Value(contextKey{})
	v, ok := v.(string)
	if !ok {
		fmt.Println("value is not a string")

		return
	}

	fmt.Println("value:", v)

	// Попытка получить значение по несуществующему ключу
	v = ctx.Value(notFoundKey{})
	if v == nil {
		fmt.Println("value is not found")
	}
}
