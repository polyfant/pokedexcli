package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		BaseStat int    `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	CaughtPokemon map[string]*Pokemon
}

func NewConfig() *config {
	return &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Hour, 5 * time.Hour),
		CaughtPokemon: make(map[string]*Pokemon),
	}
}