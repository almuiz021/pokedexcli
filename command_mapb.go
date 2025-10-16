package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *Config) error {

	if cfg.Poke.Previous != nil {
		locs, err := cfg.Poke.ListLocationPrevious()

		if err != nil {
			return err
		}
		for _, loc := range locs {
			fmt.Println(loc)
		}
	} else {
		return errors.New("you're on the first page")
	}
	return nil
}
