package weather

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// FetchDataFromAPI sends request to 3rd party weather api and returns body of response
func fetchDataFromAPI(ctx context.Context, url string) ([]byte, error) {
	weatherClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting request: %w", err)
	}

	resp, err := weatherClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting response: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("response return status %s:\nbody: %s", resp.Status, body)
	}

	fmt.Println("Data from 3rd Party API was fetched successfully")

	return body, nil
}
