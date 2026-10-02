package handlers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dannyjimenez98/weather-api.git/internal/client"
	"github.com/dannyjimenez98/weather-api.git/internal/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

type RequestParams struct {
	location string
	date1    string
	date2    string
}

func Handler(r *chi.Mux, rdb *redis.Client) {
	r.Use(middleware.StripSlashes)

	r.Get("/weather/{location}", handleWeatherData(rdb))
	r.Get("/weather/{location}/{date1}", handleWeatherData(rdb))
	r.Get("/weather/{location}/{date1}/{date2}", handleWeatherData(rdb))
}

// HandleWeatherData returns response with location's weather data for date range
func handleWeatherData(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := RequestParams{
			location: chi.URLParam(r, "location"),
			date1:    chi.URLParam(r, "date1"),
			date2:    chi.URLParam(r, "date2"),
		}

		if err := validateRequestParams(params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		const baseURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
		apiKey := env.APIKey()
		apiQuery := "?unitGroup=us&include=days&key=" + apiKey

		pathParams := params.location
		switch {
		case params.date2 != "" && params.date1 != "":
			pathParams += "/" + params.date1 + "/" + params.date2
		case params.date1 != "":
			pathParams += "/" + params.date1
		default:
		// continue with location as only path parameter
		// by default, a path without dates will return 15 day forecast for location in request
		}

		url := baseURL + pathParams + apiQuery

		ctx := r.Context()
		var body []byte

		fmt.Println("key : ", pathParams)
		cachedData, err := rdb.Get(ctx, pathParams).Bytes()
		if err != nil {
			body = client.FetchWeatherData(w, url)
			rdb.Set(ctx, pathParams, body, time.Hour*12)
		} else {
			fmt.Println("key was found!\nreturning cached data...")
			body = cachedData
		 }

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)

		keys, _ := rdb.Keys(ctx, "*").Result()
		fmt.Println("Currently cached keys: ", keys)
		fmt.Println()
	}
}

func validateRequestParams(params RequestParams) error {
	if strings.TrimSpace(params.location) == "" {
		return fmt.Errorf("location parameter is required but value was not provided")
	}

	if strings.TrimSpace(params.date1) == "" && params.date2 != "" {
		return fmt.Errorf("date2 cannot be passed in the request without date1")
	}

	return nil
}
