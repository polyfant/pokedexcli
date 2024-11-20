package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
