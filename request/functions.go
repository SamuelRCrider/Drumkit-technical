package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Create a new Motive API client
func newMotiveAPI(baseURL, apiKey string) *MotiveAPI {
	return &MotiveAPI{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Fetch a specific component from the Motive API
func (m *MotiveAPI) fetchComponent(componentType string, access_token string) (interface{}, error) {

	url := fmt.Sprintf("%s/%s", m.BaseURL, componentType)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+access_token)

	res, err := m.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", res.StatusCode)
	}

	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// Sends data to the callback URL
func sendCallback(callbackURL string, component FleetComponent) error {
	payload, err := json.Marshal(component)
	if err != nil {
		return err
	}

	res, err := http.Post(callbackURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("callback request failed with status code: %d", res.StatusCode)
	}

	return nil
}

// Fetch and send a callback for a specific component
func processComponent(api *MotiveAPI, componentType, callbackURL string, access_token string) error {

	// Fetch
	data, err := api.fetchComponent(componentType, access_token)
	if err != nil {
		return fmt.Errorf("error fetching %s: %v", componentType, err)
	}

	component := FleetComponent{
		Type: componentType,
		Data: data,
	}

	// Send
	if err := sendCallback(callbackURL, component); err != nil {
		return fmt.Errorf("error sending %s callback: %v", componentType, err)
	}

	return nil
}
