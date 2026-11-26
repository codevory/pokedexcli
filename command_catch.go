package main

import (
	"context"
	"fmt"
	"math/rand"
)

func catch(cfg *config) error {
	if cfg.entity == "" {
		return fmt.Errorf("usage: catch <pokemon name>")
	}

	pokemon, err,statusCode := cfg.pokeapiClient.CatchPokemon(context.Background(), cfg.entity)
	if err != nil && statusCode == 404 {
		return fmt.Errorf("i can't see %s",cfg.entity)
	} else if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchChance := 100 - pokemon.Experience
	if catchChance < 1 {
		catchChance = 1
	} else if catchChance > 100 {
		catchChance = 100
	}

	if rand.Intn(100) < catchChance {
		cfg.pokeapiClient.AddPokemon(pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
