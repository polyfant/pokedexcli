package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func getStartupMessage() string {
	title := `
    ██████╗  ██████╗ ██╗  ██╗███████╗██████╗ ███████╗██╗  ██╗
    ██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝██╔══██╗██╔════╝╚██╗██╔╝
    ██████╔╝██║   ██║█████╔╝ █████╗  ██║  ██║█████╗   ╚███╔╝ 
    ██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║  ██║██╔══╝   ██╔██╗ 
    ██║     ╚██████╔╝██║  ██╗███████╗██████╔╝███████╗██╔╝ ██╗
    ╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═════╝ ╚══════╝╚═╝  ╚═╝
	
	╔══════════════════════════════════════════════════════════╗
	║  ポケモン  P O K E D E X  図鑑  V1.0                    ║
	║  ---------------------------------------------          ║
	║  [ PRESS HELP TO START YOUR POKEMON JOURNEY! ]          ║
	╚══════════════════════════════════════════════════════════╝

	Available commands:
	⬢ help     ⬢ map      ⬢ explore
	⬢ catch    ⬢ inspect  ⬢ exit
	`
	rainbow := []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgYellow),
		color.New(color.FgGreen),
		color.New(color.FgCyan),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
	}

	var result string
	lines := strings.Split(title, "\n")
	for i, line := range lines {
		if i < 6 { // Only color the ASCII art title
			result += rainbow[i].Sprintf("%s\n", line)
		} else {
			result += line + "\n"
		}
	}

	return result + "\nWelcome to the retro Pokedex! Type 'help' to get started."
}

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
			callback: func(cfg *config, args []string) error {
				return commandHelp(cfg, args...)
			},
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas",
			callback: func(cfg *config, args []string) error {
				return commandMap(cfg, args...)
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(cfg *config, args []string) error {
				return commandExit(cfg, args...)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback: func(cfg *config, args []string) error {
				return commandExplore(cfg, args...)
			},
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback: func(cfg *config, args []string) error {
				return commandCatch(cfg, args...)
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback: func(cfg *config, args []string) error {
				return commandInspect(cfg, args...)
			},
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	prompt := color.New(color.FgHiCyan, color.Bold)
	errorMsg := color.New(color.FgRed)

	fmt.Print(getStartupMessage())

	for {
		prompt.Print("\nポケモン > ")
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
				errorMsg.Println("❌", err)
			}
			continue
		} else {
			errorMsg.Println("❌ Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
