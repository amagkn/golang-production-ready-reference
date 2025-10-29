package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	deadline := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go deadlineWorker(ctx, &wg)

	wg.Wait()
}

func deadlineWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Прошло 2 секунды, дедлайн истёк")
		fmt.Println("Проверяем ошибку контекста:", ctx.Err())

	case <-time.After(time.Hour):
	}
}
