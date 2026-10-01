package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dannyjimenez98/weather-api.git/internal/client"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestParams struct {
	location string
	date1    string
	date2    string
}

func Handler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)

	r.Get("/weather/{location}", handleWeatherData)
	r.Get("/weather/{location}/{date1}", handleWeatherData)
	r.Get("/weather/{location}/{date1}/{date2}", handleWeatherData)
}

// HandleWeatherData returns response with location's weather data for date range
func handleWeatherData(w http.ResponseWriter, r *http.Request) {
	params := RequestParams{
		location: chi.URLParam(r, "location"),
		date1:    chi.URLParam(r, "date1"),
		date2:    chi.URLParam(r, "date2"),
	}

	if err := validateRequestParams(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	const baseURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	apiKey := os.Getenv("API_KEY")
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

	body := client.FetchWeatherData(w, url)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	fmt.Println("url: ", url)
	fmt.Println("location: ", params.location)
	fmt.Println("date1: ", params.date1)
	fmt.Println("date2: ", params.date2)
	fmt.Println("path params: ", pathParams)
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
