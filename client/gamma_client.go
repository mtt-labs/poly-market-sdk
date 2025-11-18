package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// GammaClient is a simple HTTP client for Gamma API (read-only, no authentication required)
// Reference: https://docs.polymarket.com/developers/gamma-markets-api/overview
type GammaClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewGammaClient creates a new Gamma API client
func NewGammaClient() *GammaClient {
	return &GammaClient{
		baseURL: "https://gamma-api.polymarket.com",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get executes GET request to Gamma API
func (c *GammaClient) Get(endpoint string) ([]byte, error) {
	url := c.baseURL + endpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
