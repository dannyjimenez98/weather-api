package cache

import (
	"context"
	"log"

	"github.com/dannyjimenez98/weather-api.git/internal/env"
	"github.com/redis/go-redis/v9"
)



func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.Address(),
		Password: env.Password(),
		DB:       env.Database(),
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("could not connect to redis")
	}

	return rdb
}
