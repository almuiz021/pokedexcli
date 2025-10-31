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
		"map": {
			name:        "map",
			description: "Gets Next 20 Location-Area Names",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets Previous 20 Location-Area Names",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore locationAreaName",
			description: "Gets Pokemon in the given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch pokemonName",
			description: "Catches Pokemon and stores your pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
