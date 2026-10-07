package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dannyjimenez98/weather-api.git/internal/cache"
	"github.com/dannyjimenez98/weather-api.git/internal/env"
	"github.com/dannyjimenez98/weather-api.git/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := start(); err != nil {
		log.Fatal("http server disconnected: %v", err)
	}
}

func start() error {
	env.Load()

	// context created just for the redis ping test
	// cancel context once connection confirmed, since cache will use context from request
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rdb, err := cache.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	r := chi.NewRouter()
	handlers.RegisterRoutes(r, rdb)

	// returning instead of using log.Fatal() so that the deferred rdb.Close() can run before program exits
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}
