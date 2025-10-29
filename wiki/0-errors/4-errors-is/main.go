package main

import (
	"errors"
	"fmt"
)

func main() {
	err := Request()
	fmt.Println("request error:", err)

	fmt.Println("ErrNotFound == err:", ErrNotFound == err)

	fmt.Println("ErrNotFound == err:", errors.Is(err, ErrNotFound))
}

func Request() error {
	err := UseCase()
	if err != nil {
		return fmt.Errorf("usecase: %w", err)
	}

	return nil
}

func UseCase() error {
	err := PostgresAdapter()
	if err != nil {
		return fmt.Errorf("postgres adapter: %w", err)
	}

	err = RedisAdapter()
	if err != nil {
		return fmt.Errorf("redis adapter: %w", err)
	}

	return nil
}

var ErrNotFound = errors.New("not found")

func PostgresAdapter() error {
	return ErrNotFound
}

func RedisAdapter() error {
	return nil
}
