package main


import (
	"bufio"
	"fmt"
	"os"
)


func startRepl() {
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
			cmd.callback()
		} else {
			fmt.Println("Unknown Command")
		}
	}
}


type cliCommand struct {
	name string
	description string
	callback func() error

}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}