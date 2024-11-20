package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands(cfg *config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(cfg *config, args []string) error {
				return commandHelp(cfg, args...)
			},
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas",
			callback:    func(cfg *config, args []string) error {
				return commandMap(cfg, args...)
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(cfg *config, args []string) error {
				return commandExit(cfg, args...)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    func(cfg *config, args []string) error {
				return commandExplore(cfg, args...)
			},
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    func(cfg *config, args []string) error {
				return commandCatch(cfg, args...)
			},
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		 words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands(cfg)[commandName]
		if exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
