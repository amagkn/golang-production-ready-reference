package main

import (
	"context"
	"fmt"
	"time"
)

type ctxKey struct{}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for range 3 {
		go Controller(ctx)
	}

	time.Sleep(5 * time.Second)
}

func Controller(ctx context.Context) {
	ctx = context.WithValue(ctx, ctxKey{}, "I am Controller")

	go Handler(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Controller done!")
	}
}

func Handler(ctx context.Context) {
	ctx = context.WithValue(ctx, ctxKey{}, "I am Handler")

	go UseCase(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Handler done!")
	}
}

func UseCase(ctx context.Context) {
	fmt.Println("Value from usecase:", ctx.Value(ctxKey{}))

	select {
	case <-ctx.Done():
		fmt.Println("UseCase done!")
	}
}
