package main

import (
	"fmt"
	"os"

	"github.com/mpetkov228/pokedex/internal/pokecache"
)

func commandExit(config *Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
