package main

import (
	"time"

	"github.com/mpetkov228/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
