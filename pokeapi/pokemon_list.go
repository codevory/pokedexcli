package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListPokemons(areaName string) (pokemonShallowNames, error) {
	url := baseURL + "/location-area" + "/" + areaName

	// cache hit, serve cached data
	if cachedData, ok := c.cache.Get(url); ok {
		fmt.Printf("[CACHE HIT]  %s (%v)\n", url, time.Since(time.Now()))
		pokemons := pokemonShallowNames{}
		err := json.Unmarshal(cachedData, &pokemons)
		if err != nil {
			return pokemonShallowNames{}, err
		}

		return pokemons, nil
	}

	// cache miss, make an api request
	fmt.Printf("[HTTP GET]   %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonShallowNames{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonShallowNames{}, err
	}

	defer resp.Body.Close()
	fmt.Printf("[HTTP %d]   %s (%v)\n", resp.StatusCode, url, time.Since(time.Now()))

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonShallowNames{}, err
	}

	// add data to cache
	c.cache.Add(url, data)
	fmt.Printf("Added data to cache with key : %v", url)

	pokemons := pokemonShallowNames{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return pokemonShallowNames{}, err
	}

	return pokemons, nil
}
