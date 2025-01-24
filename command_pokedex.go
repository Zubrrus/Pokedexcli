package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("you haven't catched any pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
