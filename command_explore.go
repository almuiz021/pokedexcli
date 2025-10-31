package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("must provide with location name")
	}

	areaName := args[0]
	respLocAreas, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)

	fmt.Println("Found Pokemon: ")
	for _, pokemonDetails := range respLocAreas.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonDetails.Pokemon.Name)
	}

	return nil
}
