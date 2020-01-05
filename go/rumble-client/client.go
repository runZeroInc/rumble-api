package client

import (
	rapi "github.com/RumbleDiscovery/rumble-api/go"
	"os"
)

// Config provides overrides the API host and key
type Config struct {
	Host string
	Key  string
}

// NewClient retuns a ready to use Rumble API Client
func NewClient(c *Config) *rapi.APIClient {

	// Configure the endpoint host
	apiHost := "console.rumble.run"
	if envHost := os.Getenv("RUMBLE_API_HOST"); envHost != "" {
		apiHost = envHost
	}
	if c.Host != "" {
		apiHost = c.Host
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
	return rapi.NewAPIClient(&rapi.Configuration{
		Host:          apiHost,
		BasePath:      "/api/v1.0",
		Scheme:        "https",
		DefaultHeader: headers,
	})
}
