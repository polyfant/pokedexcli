package main
import (
	"time"
	"pokedex/internal/pokeapi"
	"fmt"
	"bufio"
	"os"
)

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		
		if input == "exit" {
			break
		}
		// Handle other commands here
	}
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: *pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
