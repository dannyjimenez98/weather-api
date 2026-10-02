package client

import (
	"fmt"
	"io"
	"net/http"
)

// FetchWeatherData sends request to 3rd party weather api and returns body of response
func FetchWeatherData(w http.ResponseWriter, url string) []byte {
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

	fmt.Println("Data from 3rd Party API was fetched successfully")

	return body
}
