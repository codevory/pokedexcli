package main 

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
  locations,err := cfg.pokeapiClient.ListLocations(cfg.NextLocationsURL)
  if err != nil {
	return err
  }

  cfg.NextLocationsURL = locations.Next
  cfg.PrevLocationsURL = locations.Previous

  for _,loc := range locations.Results {
	fmt.Println(loc.Name)
  }

  return nil;
}

func commandMapB(cfg *config) error {
	if cfg.PrevLocationsURL == nil {
      return errors.New("You're on the first page")
	}

	locations,err := cfg.pokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locations.Next 
	cfg.PrevLocationsURL = locations.Previous

	for _,loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}