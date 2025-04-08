package main

import (
	"time"

	"github.com/mpetkov228/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(2 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
