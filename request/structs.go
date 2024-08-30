package request

import "net/http"

// Motive API client
type MotiveAPI struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// Vehicles, drivers, or trailers
type FleetComponent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
