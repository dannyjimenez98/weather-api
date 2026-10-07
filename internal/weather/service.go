package weather

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const cacheTTL = time.Hour

func GetWeatherResponse(ctx context.Context, params *RequestParams, rdb *redis.Client) ([]byte, error) {
	requestURL := params.buildRequestURL()
	cacheKey := params.generateKey()

	var responseBody []byte
	cachedData, err := rdb.Get(ctx, cacheKey).Bytes()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	// cache miss
	if errors.Is(err, redis.Nil) {
		responseBody, err = fetchDataFromAPI(ctx, requestURL)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve weather data: %w", err)
		}
		rdb.Set(ctx, cacheKey, responseBody, cacheTTL)
	} else {
		fmt.Println("key was found!\nreturning cached data...")
		responseBody = cachedData
	}
	return responseBody, nil
}
