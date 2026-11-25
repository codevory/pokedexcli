package pokeapi

import (
	"github.com/codevory/pokedexcli/pokecache"
	"net/http"
	"time"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
	pokedex    map[string]PokemonCaught
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokedex: make(map[string]PokemonCaught),
	}
}

func (c *Client) AddPokemon(pokemon PokemonCaught) {
	c.pokedex[pokemon.Name] = pokemon
}
