package main

import (
	"context"
	"fmt"
)

func main() {
	{ // Контекст по умолчанию
		ctx := context.Background()

		fmt.Println("ctx.Err():", ctx.Err())       // <nil>
		fmt.Println("ctx.Done():", ctx.Done())     // <nil>
		fmt.Println("ctx.Value", ctx.Value("key")) // <nil>
		fmt.Println(ctx.Deadline())                // 0001-01-01 00:00:00 +0000 UTC false
	}

	{ // Временный контекст
		ctx := context.TODO()
		_ = ctx
	}
}
