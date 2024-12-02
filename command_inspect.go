package main

import (
	"errors"
	"fmt"
	"strings"
	"github.com/fatih/color"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s yet", pokemonName)
	}

	titleColor := color.New(color.FgHiYellow, color.Bold)
	statColor := color.New(color.FgHiCyan)
	
	fmt.Println("╔" + strings.Repeat("═", 50) + "╗")
	titleColor.Printf("║ %s's Data\n", strings.ToUpper(pokemon.Name))
	fmt.Println("╠" + strings.Repeat("═", 50) + "╣")
	
	statColor.Printf("║ Height: %d\n", pokemon.Height)
	statColor.Printf("║ Weight: %d\n", pokemon.Weight)
	
	fmt.Println("║")
	titleColor.Println("║ Stats:")
	for _, stat := range pokemon.Stats {
		barLength := stat.BaseStat / 10
		statBar := strings.Repeat("█", barLength) + strings.Repeat("░", 10-barLength)
		statColor.Printf("║  %s: %s %d\n", 
			padRight(stat.Stat.Name, 12),
			statBar,
			stat.BaseStat)
	}
	
	fmt.Println("║")
	titleColor.Println("║ Types:")
	for _, typeInfo := range pokemon.Types {
		statColor.Printf("║  ⬢ %s\n", typeInfo.Type.Name)
	}
	
	fmt.Println("╚" + strings.Repeat("═", 50) + "╝")
	return nil
}

func padRight(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}
