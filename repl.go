package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/almuiz021/pokedexcli/internal/pokeapi"
)

type Config struct {
	Poke *pokeapi.Client
}

func startRepl(cfg *Config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		inputCmd := reader.Text()

		if len(inputCmd) == 0 {
			continue
		}

		cmd, exists := getCommands()[inputCmd]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
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
			description: "Provides the name of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Provides the name of previous Locations",
			callback:    commandMapb,
		},
	}
}
