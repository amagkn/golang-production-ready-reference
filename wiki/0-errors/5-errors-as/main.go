package main

import (
	"errors"
	"fmt"
)

func main() {
	err := Request()
	fmt.Println("request error:", err)

	var pgErr PostgresError
	if errors.As(err, &pgErr) {
		fmt.Println("postgres error code:", pgErr.Code)
	}
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

func PostgresAdapter() error {
	return PostgresError{Context: "connection refused", Code: 42}
}

func RedisAdapter() error {
	return nil
}
