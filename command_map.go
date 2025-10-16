package main

import "fmt"

func commandMap(cfg *Config) error {

	locs, err := cfg.Poke.ListLocationNext()
	if err != nil {
		return err
	}
	for _, loc := range locs {
		fmt.Println(loc)
	}
	return nil

}
