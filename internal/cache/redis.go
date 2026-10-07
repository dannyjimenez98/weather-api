package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/dannyjimenez98/weather-api.git/internal/env"
	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:                  env.Address(),
		Password:              env.Password(),
		DB:                    env.Database(),
		ContextTimeoutEnabled: true,
		DialTimeout:           time.Second,
		ReadTimeout:           time.Second,
		WriteTimeout:          time.Second,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		rdb.Close()
		return nil, fmt.Errorf("could not connect to redis: %w", err)
	}

	return rdb, nil
}
