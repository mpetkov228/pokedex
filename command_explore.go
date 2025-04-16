package main

import "fmt"

func explore(config *Config, args ...string) error {
	location := args[0]
	data, err := config.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		name := encounter.Pokemon.Name
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
