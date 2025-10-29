package main

import "fmt"

type PostgresError struct {
	Code    int
	Context string
}

func (e PostgresError) Error() string {
	return fmt.Sprintf("postgres error: code %d: %s", e.Code, e.Context)
}

type RedisError struct {
	Message string
}

func (e RedisError) Error() string {
	return fmt.Sprintf("redis error: %s", e.Message)
}
