package database

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/TimeleapLabs/unchained/internal/ent"
)

type Database interface {
	GetConnection() *ent.Client
	HealthCheck(ctx context.Context) bool
}

type IRedisDatabase interface {
	GetConnection() *redis.Client
	HealthCheck(ctx context.Context) bool
}
