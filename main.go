package main

import (
	"time"

	"github.com/Zubrrus/Pokedexcli/internal/pokeapi"
)

// minBaseExp: 36 Sunkern
// maxBaseExp: 635 Blissey

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
