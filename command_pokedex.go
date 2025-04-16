package main

import "fmt"

func pokedex(config *Config, args ...string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key := range config.pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
