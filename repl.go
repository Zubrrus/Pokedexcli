package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zubrrus/Pokedexcli/internal/pokeapi"
)

type config struct {
	caughtPokemon    map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		params := []string{}
		if len(words) > 1 {
			params = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, params...)
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
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Displays all pokemons presented on specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokenon_name>",
			description: "Catching specified pokemon with certain chance",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokenon_name>",
			description: "Displays pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all catched pokemons",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
