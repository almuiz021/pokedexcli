package main

import (
	"strings"

	"github.com/almuiz021/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	stringArr := strings.Fields(lowerCased)
	return stringArr
}

func main() {
	cfg := &Config{Poke: &pokeapi.Client{}}
	startRepl(cfg)
}
