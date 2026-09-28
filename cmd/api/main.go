package main

import (
	"log"
	"net/http"

	"github.com/dannyjimenez98/weather-api.git/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}

	mux := http.NewServeMux()

	// get location's weather data for today
	mux.HandleFunc("/{location}/today", api.GetWeatherToday)

	// get location's weather data for today
	mux.HandleFunc("/{location}/{date}", api.GetWeatherOnDate)

	// get location's weather data for today
	mux.HandleFunc("/{location}/{startDate}/{endDate}", api.GetWeatherBetweenDates)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
