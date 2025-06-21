package main

import (
	"time"

	"github.com/gdbeltran/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := &cliConfig{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
