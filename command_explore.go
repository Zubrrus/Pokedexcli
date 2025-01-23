package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("location to explore is not specified")
	}

	locationName := params[0]

	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounters := range pokemonsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounters.Pokemon.Name)
	}
	return nil
}
