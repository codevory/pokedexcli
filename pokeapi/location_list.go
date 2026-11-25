package pokeapi

import (
	"context"
)

func (c *Client) ListLocations(pageUrl *string,ctx context.Context) (RespShallowLocations, error,int) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

  locations,err,StatusCode := fetchJSON[RespShallowLocations](ctx,c,url)
  
  if err != nil {
	return RespShallowLocations{},err,StatusCode
  }

	return locations, nil,StatusCode
}
