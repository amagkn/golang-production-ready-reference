package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		fmt.Println("Получен сигнал:", <-sig)

		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go cancelWorker(ctx, &wg)

	wg.Wait()
}

func cancelWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Проверяем ошибку контекста:", ctx.Err())

	case <-time.After(time.Hour):
	}
}
