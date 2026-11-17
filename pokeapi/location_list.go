package pokeapi

import (
	"io"
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string)(RespShallowLocations,error){
  url := baseURL + "/location-area"
  if pageUrl != nil {
	url = * pageUrl
  }

  req,err := http.NewRequest("GET",url,nil)
  if err != nil {
	return RespShallowLocations{},err
  }

  resp,err := c.httpClient.Do(req)
  if err != nil {
	return RespShallowLocations{},err
  }

  defer resp.Body.Close()

  data,err := io.ReadAll(resp.Body)
  if err != nil {
	return RespShallowLocations{},err
  }

  locations := RespShallowLocations{}
  err = json.Unmarshal(data,&locations)
  if err != nil {
	return RespShallowLocations{},err
  }

  return locations,nil
}