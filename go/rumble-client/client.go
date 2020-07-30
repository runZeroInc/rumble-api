package client

import (
	"os"

	rapi "github.com/RumbleDiscovery/rumble-api/go"
)

// Config provides overrides the API host and key
type Config struct {
	URL string
	Key string
}

// NewClient retuns a ready to use Rumble API Client
func NewClient(c *Config) *rapi.APIClient {

	// Configure the endpoint host
	apiURL := "https://console.rumble.run/api/v1.0"
	if envHost := os.Getenv("RUMBLE_API_URL"); envHost != "" {
		apiURL = envHost
	}
	if c.URL != "" {
		apiURL = c.URL
	}

	// Configure the authorization header
	headers := make(map[string]string)
	apiKey := os.Getenv("RUMBLE_API_KEY")
	if c.Key != "" {
		apiKey = c.Key
	}
	if apiKey != "" {
		headers["Authorization"] = "Bearer " + apiKey
	}

	// Create the client
	config := rapi.NewConfiguration()
	config.DefaultHeader = headers
	config.Servers = rapi.ServerConfigurations{
		{
			URL:         apiURL,
			Description: "Rumble Console",
		},
	}

	return rapi.NewAPIClient(config)

	/*
		return rapi.NewAPIClient(&rapi.Configuration{
			Host:          apiHost,
			BasePath:      "/api/v1.0",
			Scheme:        "https",
			DefaultHeader: headers,
		})
	*/
}
