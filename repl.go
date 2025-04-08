package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mpetkov228/pokedex/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]
		if ok {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command.")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List world locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous world locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
