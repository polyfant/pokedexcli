package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	cfg.CaughtPokemon[pokemon.Name] = &Pokemon{
		Name:   pokemon.Name,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
		Stats:  convertStats(pokemon.Stats),
		Types:  convertTypes(pokemon.Types),
	}
	return nil
}

func convertStats(apiStats []struct {
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}) []struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
} {
	stats := make([]struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	}, len(apiStats))

	for i, s := range apiStats {
		stats[i].BaseStat = s.BaseStat
		stats[i].Stat.Name = s.Stat.Name
	}
	return stats
}

func convertTypes(apiTypes []struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}) []struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
} {
	types := make([]struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	}, len(apiTypes))

	for i, t := range apiTypes {
		types[i].Type.Name = t.Type.Name
	}
	return types
}
