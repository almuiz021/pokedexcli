package main

import (
	"strings"
	"fmt"
)


func cleanInput(text string) []string{

	lowerCased := strings.ToLower(text)
	stringArr := strings.Fields(lowerCased)
	return stringArr
}


func main(){
	fmt.Println("testing")
}
