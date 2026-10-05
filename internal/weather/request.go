package weather

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/dannyjimenez98/weather-api.git/internal/env"
)

type RequestParams struct {
	Location string
	Date1    string
	Date2    string
}

func (rp *RequestParams) buildRequestURL() string {
	const baseURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/"
	query := url.Values{
		"unitGroup": {"us"},
		"include":   {"days"},
		"key":       {env.APIKey()},
	}

	path := url.PathEscape(rp.Location)
	if rp.Date1 != "" {
		path += "/" + url.PathEscape(rp.Date1)
	}
	if rp.Date2 != "" {
		path += "/" + url.PathEscape(rp.Date2)
	}

	return baseURL + path + "?" + query.Encode()
}

// GenerateKey builds a Redis cache key from the API request parameters.
func (rp *RequestParams) generateKey() string {
	weatherCacheKey := rp.Location + ":" + rp.Date1 + ":" + rp.Date2

	return weatherCacheKey
}

func (rp *RequestParams) Validate() error {
	if strings.TrimSpace(rp.Location) == "" {
		return fmt.Errorf("location parameter is required but value was not provided")
	}

	if strings.TrimSpace(rp.Date1) == "" && rp.Date2 != "" {
		return fmt.Errorf("date2 cannot be passed in the request without date1")
	}

	return nil
}
