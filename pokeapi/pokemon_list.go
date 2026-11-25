package pokeapi

import (
	"context"
)

func (c *Client) ListPokemons(areaName string,ctx context.Context) (pokemonShallowNames, error,int) {
	url := baseURL + "/location-area" + "/" + areaName

	res,err,statusCode := fetchJSON[pokemonShallowNames](ctx,c,url)
	if err != nil {
		return pokemonShallowNames{},err,statusCode
	}

	return res,nil,statusCode
}
