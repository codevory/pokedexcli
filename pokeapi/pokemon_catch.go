package pokeapi

import "context"

func (c *Client) CatchPokemon(ctx context.Context, name string) (PokemonCaught, error,int) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	res, err,statusCode := fetchJSON[PokemonCaught](ctx, c, url)
	if err != nil {
		return PokemonCaught{}, err,statusCode
	}

	return res, nil,statusCode
}
