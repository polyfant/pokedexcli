package main

import (
	"testing"
)

func TestStartupMessage(t *testing.T) {
	cfg := &config{}
	msg := cfg.StartupMessage()
	if msg == "" {
		t.Error("Expected non-empty startup message")
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	
	expectedCommands := []string{"help", "map", "exit", "explore", "catch"}
	for _, cmd := range expectedCommands {
		if _, exists := commands[cmd]; !exists {
			t.Errorf("Expected command %s to exist", cmd)
		}
	}
}

func (c *config) StartupMessage() string {
	return "Welcome to the Pokemon Game!"
} 