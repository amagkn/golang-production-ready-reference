package redis

import (
	"time"

	"github.com/amagkn/golang-production-ready-reference/pkg/redis"
)

const (
	idempotencyPrefix = "my-app:idempotency:"
	ttl               = time.Hour
)

type Redis struct {
	redis *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		redis: client,
	}
}
