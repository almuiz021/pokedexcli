package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	resLocs, err := cfg.pokeapiClient.FetchLocations(cfg.Next)
	if err != nil {
		return err
	}

	locAreas := resLocs.Results

	cfg.Next = resLocs.Next
	cfg.Previous = resLocs.Previous

	for _, loc := range locAreas {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	resLocs, err := cfg.pokeapiClient.FetchLocations(cfg.Previous)

	if err != nil {
		return err
	}

	locAreas := resLocs.Results

	cfg.Next = resLocs.Next
	cfg.Previous = resLocs.Previous

	for _, loc := range locAreas {
		fmt.Println(loc.Name)
	}

	return nil
}
