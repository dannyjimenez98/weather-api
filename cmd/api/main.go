package main

import (
	"log"
	"net/http"

	"github.com/dannyjimenez98/weather-api.git/internal/api"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env from the working directory (run from the project root): %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/weather/data", api.GetWeatherData)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
