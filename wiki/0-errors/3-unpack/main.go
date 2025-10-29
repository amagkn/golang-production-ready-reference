package main

import (
	"errors"
	"fmt"
)

func main() {
	err := Request()
	fmt.Println("request error:", err)

	err = errors.Unwrap(err)
	fmt.Println("unwrap:", err)

	err = errors.Unwrap(err)
	fmt.Println("unwrap:", err)

	err = errors.Unwrap(err)
	fmt.Println("unwrap:", err)
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

var ErrNotFound = fmt.Errorf("not found")

func PostgresAdapter() error {
	return ErrNotFound
}

func RedisAdapter() error {
	return nil
}
