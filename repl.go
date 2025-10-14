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
		
		inputText := reader.Text()

		if len(inputText) == 0 {
			continue
		}

		firstWord := cleanInput(inputText)[0]
		
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}