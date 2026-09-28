package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// GetWeatherData sends a request to the 3rd party weather api
// and writes the returned data to the response of this endpoint
func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	location := "10121" // NYC zipcode

	apiKey := os.Getenv("API_KEY")
	todaysDate := time.Now().Format("2006-01-02")

	// endpoint returns weather data for today only
	url := fmt.Sprintf("%s%s/%s?unitGroup=us&include=days&key=%s", baseURL, location, todaysDate, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error getting response: %v\n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v\n", err)
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
