package main

import (
	"fmt"
	"math/rand"
)

func catch(config *Config, args ...string) error {
	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catchChance := getChance(pokemon.BaseExperience)
	n := rand.Intn(100)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if n <= catchChance {
		config.pokedex[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func getChance(exp int) int {
	catchChance := 90 - (exp * 80 / 300)
	if catchChance < 10 {
		return 10
	}
	return catchChance
}
