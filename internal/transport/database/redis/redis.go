package redis

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/redis/go-redis/v9"
)

type connection struct {
	db *redis.Client
}

func (c *connection) HealthCheck(ctx context.Context) bool {
	if c.db != nil {
		return true
	}

	status := c.db.Ping(ctx)
	if status.Err() != nil {
		utils.Logger.Error("HealthCheck: " + status.Err().Error())
		return false
	}

	return true
}

func (c *connection) GetConnection() *redis.Client {
	if c.db != nil {
		return c.db
	}

	utils.Logger.Info("Connecting to Redis")
	var err error
	opts, err := redis.ParseURL(config.App.Redis.Dsn)
	if err != nil {
		panic(err)
	}

	c.db = redis.NewClient(opts)

	return c.db
}

func New() database.IRedisDatabase {
	return &connection{}
}
