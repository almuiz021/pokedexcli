package main

import (
	"time"

	"github.com/almuiz021/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// config is on heap and cfg is on stack [ 48bytes and 8bytes ]
	// cfg points to config
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
