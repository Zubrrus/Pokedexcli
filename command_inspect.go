package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("pokemon name not specified")
	}

	pokemonName := params[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pokemon_stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokemon_stat.Stat.Name, pokemon_stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemon_type := range pokemon.Types {
		fmt.Printf("  -%s\n", pokemon_type.Type.Name)
	}

	return nil
}
