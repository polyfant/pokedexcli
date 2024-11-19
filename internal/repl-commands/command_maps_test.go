package commands

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "MAP   ",
			expected: []string{"map"},
		},
		{
			input:    "  map next  ",
			expected: []string{"map", "next"},
		},
		{
			input:    "MAP BACK   ",
			expected: []string{"map", "back"},
		},
	}

	for _, c := range cases {
		got := cleanInput(c.input)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("cleanInput(%s) = %v; expected %v", c.input, got, c.expected)
		}
	}
}

func TestGetCommandName(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "map",
			expected: "map",
		},
		{
			input:    "map next",
			expected: "map",
		},
	}
	for _, c := range cases {
		got := cleanInput(c.input)[0]
		if got != c.expected {
			t.Errorf("getCommandName(%s) = %s; expected %s", 
				c.input, got, c.expected)
		}
	}
} 