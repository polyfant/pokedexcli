package pokeapi

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	// Check if the cache is properly initialized
	if client.cache.cache == nil {
		t.Error("Expected cache map to be initialized")
	}
}

func TestListLocations(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	resp, err := client.ListLocations(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(resp.Results) == 0 {
		t.Error("Expected at least one location")
	}
}

func TestGetLocation(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	resp, err := client.GetLocation("test-location")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp.Name == "" {
		t.Error("Expected location to have a name")
	}
}