package pokeapi

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	cache map[string][]byte
	httpClient http.Client
}

// NewClient -
func NewClient(cacheInterval time.Duration, retryInterval time.Duration) *Client {
	return &Client{
		cache: make(map[string][]byte),
		httpClient: http.Client{
			Timeout: retryInterval,
		},
	}
}
