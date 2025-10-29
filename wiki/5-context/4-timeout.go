package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	now := time.Now()

	for range 10 {
		wg.Add(1)
		go timeoutWorker(ctx, &wg)
	}

	wg.Wait()

	fmt.Println("Время выполнения:", time.Since(now))
}

func timeoutWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Проверяем ошибку контекста:", ctx.Err())

	case <-time.After(time.Hour):
	}
}
