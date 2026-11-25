package main

import (
	"fmt"
	"context"
)

func explore(cfg *config) error {
	area, err,statusCode := cfg.pokeapiClient.ListPokemons(cfg.entity,context.Background())
	if err != nil && statusCode == 404 {
		return fmt.Errorf("%s,doesn't exist on map",cfg.entity)
	} else if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", cfg.entity)
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil
}
