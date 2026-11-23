package main

import (
	"fmt"
)

func explore(cfg *config) error {
	area, err := cfg.pokeapiClient.ListPokemons(cfg.areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", cfg.areaName)
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}
