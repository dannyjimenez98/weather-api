package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// getWeatherData sends request to 3rd party weather api and returns body of response
func getWeatherData(w http.ResponseWriter, url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error getting response: "+err.Error(), http.StatusBadGateway)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error reading response body: "+err.Error(), http.StatusBadGateway)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		errorMessage := fmt.Sprintf("response failed: %v\nbody: %v", err, body)
		http.Error(w, errorMessage, resp.StatusCode)
		return nil
	}

	return body
}

// GetWeatherToday returns response with location's weather data for today
func GetWeatherToday(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	location := r.PathValue("location")
	apiKey := os.Getenv("API_KEY")

	url := baseURL + location + "/today?unitGroup=us&include=days&key=" + apiKey

	body := getWeatherData(w, url)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetWeatherOnDate returns response with location's weather data for requested date
func GetWeatherOnDate(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	location := r.PathValue("location")
	date := r.PathValue("date")

	apiKey := os.Getenv("API_KEY")

	url := baseURL + location + "/" + date + "?unitGroup=us&include=days&key=" + apiKey

	body := getWeatherData(w, url)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetWeatherBetweenDates returns response with location's weather data for date range
func GetWeatherBetweenDates(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	location := r.PathValue("location")
	startDate := r.PathValue("startDate")
	endDate := r.PathValue("endDate")

	apiKey := os.Getenv("API_KEY")

	url := baseURL + location + "/" + startDate + "/" + endDate + "?unitGroup=us&include=days&key=" + apiKey

	body := getWeatherData(w, url)

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
