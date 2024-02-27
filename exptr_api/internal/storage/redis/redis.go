package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func MustLoad() error {
	const op = "redis.MustLoad"

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	RedisClient = client
	return nil
}
