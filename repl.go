package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gdbeltran/pokedexcli/internal/pokeapi"
)

type cliConfig struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	areaName         *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
}

func startRepl(cfg *cliConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		commandParameter := input[1]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lowerWords := strings.ToLower(text)
	words := strings.Fields(lowerWords)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next page of locations",
			callback:    commandMapN,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of locations",
			callback:    commandMapP,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
	}
}
