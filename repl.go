package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/almuiz021/pokedexcli/internal/pokeapi"
)

// APP STATE
type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Pokedex       map[string]pokeapi.RespPokemon
}

type Pokemon struct {
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		inputText := cleanInput(reader.Text())

		if len(inputText) == 0 {
			continue
		}

		inputCmd := inputText[0]
		args := []string{}
		if len(inputText) > 1 {
			args = inputText[1:]
		}

		cmd, exists := getCommands()[inputCmd]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	stringArr := strings.Fields(lowerCased)
	return stringArr
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
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get all the caught pokemon names",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
