package main 

import "fmt"

func pokedex(cfg *config)error {
  store,err := cfg.pokeapiClient.GetPokemonsCaughtList()

  fmt.Println("Your Pokedex:")

  if err != nil || len(store) < 1{
	return nil
  }

  for _,pok := range store {
	fmt.Println(pok.Name)
  }

  return nil
}