package main

import (
	"fmt"

	"github.com/mpetkov228/pokedex/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
