package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config, args ...string) error {
	url := config.Next
	data, err := config.pokeapiClient.GetLocations(url)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(config *Config, args ...string) error {
	url := config.Previous
	if url == nil {
		return errors.New("you're on the first page")
	}

	data, err := config.pokeapiClient.GetLocations(url)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}
