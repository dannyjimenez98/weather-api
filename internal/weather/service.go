package weather

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetWeatherResponse(ctx context.Context, params *RequestParams, rdb *redis.Client) ([]byte, error) {
	requestURL := params.buildRequestURL()
	cacheKey := params.generateKey()

	var responseBody []byte
	cachedData, err := rdb.Get(ctx, cacheKey).Bytes()
	if err != nil {
		responseBody, err = fetchDataFromAPI(ctx, requestURL)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve weather data: %w", err)
		}
		rdb.Set(ctx, cacheKey, responseBody, time.Hour)
	} else {
		fmt.Println("key was found!\nreturning cached data...")
		responseBody = cachedData
	}
	return responseBody, nil
}
