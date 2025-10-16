package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for cmdName, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmdName, cmd.description)
	}
	// os.Exit(0)
	return nil
}
