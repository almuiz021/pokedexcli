package main

import (
	"errors"
	"fmt"

	"github.com/almuiz021/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("need name of  pokemon to inspect")
	}

	pokemonName := args[0]

	details, exists := cfg.Pokedex[pokemonName]

	if exists {
		pokemonDetails(details)
	} else {
		return errors.New("you have not caught that pokemon")
	}

	return nil
}

func pokemonDetails(pokemon pokeapi.RespPokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: ")

	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}
}
