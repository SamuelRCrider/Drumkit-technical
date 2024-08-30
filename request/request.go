package request

import (
	"fmt"
	"os"
)

func HandleRequest(access_token string) error {
	const (
		Drivers      = "v1/users"
		Vehicles     = "v1/vehicles"
		Trailers     = "v1/assets"
		Callback_URL = "https://eomskv62ag85o1l.m.pipedream.net"
	)

	api_key := os.Getenv("MOTIVE_API_KEY")

	api := newMotiveAPI("https://api.gomotive.com", api_key)

	// Process drivers
	err := processComponent(api, Drivers, Callback_URL, access_token)
	if err != nil {
		return fmt.Errorf("error processing drivers: %v", err)
	}

	// Process vehicles
	err = processComponent(api, Vehicles, Callback_URL, access_token)
	if err != nil {
		return fmt.Errorf("error processing vehicles: %v", err)
	}

	// Process trailers
	err = processComponent(api, Trailers, Callback_URL, access_token)
	if err != nil {
		return fmt.Errorf("error processing trailers: %v", err)
	}

	return nil
}
