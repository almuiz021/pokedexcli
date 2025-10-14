package main

import (
	"strings"
)


func cleanInput(text string) []string{
	lowerCased := strings.ToLower(text)
	stringArr := strings.Fields(lowerCased)
	return stringArr
}


func main(){
	startRepl()
}
