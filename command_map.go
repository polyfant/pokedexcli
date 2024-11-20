package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationAreas, err := cfg.pokeapiClient.ListLocations(nil)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
} 