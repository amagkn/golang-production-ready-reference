package main

import "fmt"

func main() {
	err := Request()
	fmt.Println("request error:", err)
}

func Request() error {
	err := UseCase()
	if err != nil {
		return err
	}

	return nil
}

func UseCase() error {
	err := PostgresAdapter()
	if err != nil {
		return err
	}

	err = RedisAdapter()
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrPostgresAdapterNotFound = fmt.Errorf("postgres adapter not found")
	ErrRedisAdapterNotFound    = fmt.Errorf("redis adapter not found")
)

func PostgresAdapter() error {
	return ErrPostgresAdapterNotFound
}

func RedisAdapter() error {
	return nil
}
