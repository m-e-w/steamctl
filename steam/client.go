package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	APIKey string
	Debug  bool
	http   *http.Client
}

// NewClient creates a new Steam API client using the provided API key
func NewClient(apiKey string, debug bool) *Client {
	return &Client{
		APIKey: apiKey,
		Debug:  debug,
		http: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// get performs a GET request against a Steam Web API endpoint and decodes the JSON response.
func (c *Client) get(endpoint string, params url.Values, v any) error {
	params.Set("key", c.APIKey)
	params.Set("format", "json")

	fullURL := endpoint + "?" + params.Encode()

	if c.Debug {
		fmt.Printf("[DEBUG] GET %s\n", endpoint)
	}

	resp, err := c.http.Get(fullURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if c.Debug {
		fmt.Printf("[DEBUG] STATUS CODE: %d\n", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("steam API request failed: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
