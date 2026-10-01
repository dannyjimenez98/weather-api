package main

import (
	"log"
	"net/http"

	"github.com/dannyjimenez98/weather-api.git/internal/handlers"
	"github.com/go-chi/chi"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}




	// http routing
	r := chi.NewRouter()
	handlers.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
