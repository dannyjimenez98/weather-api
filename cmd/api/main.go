package main

import (
	"log"
	"net/http"

	"github.com/dannyjimenez98/weather-api.git/internal/cache"
	"github.com/dannyjimenez98/weather-api.git/internal/env"
	"github.com/dannyjimenez98/weather-api.git/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	env.Load()

	rdb := cache.Connect()
	defer rdb.Close()

	// http routing
	r := chi.NewRouter()
	handlers.RegisterRoutes(r, rdb)

	log.Fatal(http.ListenAndServe(":8080", r))
}
