package main

import (
	"testing"
)

func TestStartupMessage(t *testing.T) {
	msg := getStartupMessage()
	if msg == "" {
		t.Error("Expected non-empty startup message")
	}
}

func TestGetCommands(t *testing.T) {
	cfg := NewConfig()
	commands := getCommands(cfg)
	
	expectedCommands := []string{"help", "map", "exit", "explore", "catch"}
	for _, cmd := range expectedCommands {
		if _, exists := commands[cmd]; !exists {
			t.Errorf("Expected command %s to exist", cmd)
		}
	}
} 