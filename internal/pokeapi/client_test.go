package pokeapi

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	if client == nil {
		t.Error("Expected non-nil client")
	}
}

func TestListLocations(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	resp, err := client.ListLocations(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Error("Expected non-nil response")
	}
}

func TestGetLocation(t *testing.T) {
	client := NewClient(time.Hour, time.Hour)
	resp, err := client.GetLocation("test-location")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Error("Expected non-nil response")
	}
} 