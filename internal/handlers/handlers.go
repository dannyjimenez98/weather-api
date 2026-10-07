package handlers

import (
	"net/http"

	"github.com/dannyjimenez98/weather-api.git/internal/weather"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(r *chi.Mux, rdb *redis.Client) {
	r.Use(middleware.StripSlashes)

	r.Get("/weather/{location}", handleData(rdb))
	r.Get("/weather/{location}/{date1}", handleData(rdb))
	r.Get("/weather/{location}/{date1}/{date2}", handleData(rdb))
}

// HandleData returns response with location's weather data for date range
func handleData(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := weather.RequestParams{
			Location: chi.URLParam(r, "location"),
			Date1:    chi.URLParam(r, "date1"),
			Date2:    chi.URLParam(r, "date2"),
		}

		if err := params.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		responseBody, err := weather.GetWeatherResponse(ctx, &params, rdb)
		if err != nil {
			http.Error(w, "could not retrieve weather data", http.StatusBadGateway)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}
