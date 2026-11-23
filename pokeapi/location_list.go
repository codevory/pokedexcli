package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// cache hit, serve cached data
	if cachedData, ok := c.cache.Get(url); ok {
		fmt.Printf("[CACHE HIT]  %s (%v)\n", url, time.Since(time.Now()))
		locations := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locations)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locations, nil
	}

	// cache miss, make an api request
	fmt.Printf("[HTTP GET]   %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()
	fmt.Printf("[HTTP %d]   %s (%v)\n", resp.StatusCode, url, time.Since(time.Now()))

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// add data to cache
	c.cache.Add(url, data)
	fmt.Printf("Added data to cache with key : %v", url)

	locations := RespShallowLocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locations, nil
}
