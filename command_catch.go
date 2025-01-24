package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("pokemon name not specified")
	}

	pokemonName := params[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	diceRoll := rand.Intn(700)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	if diceRoll < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
