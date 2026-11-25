package pokeapi 

import "fmt"

func(c *Client) GetPokemonsCaughtList() (map[string]PokemonCaught,error){
	store := c.pokedex 
	if len(store) < 1 {
		return nil,fmt.Errorf("No pokemons caught")
	}

	return store,nil
}