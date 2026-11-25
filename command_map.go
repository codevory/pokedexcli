package main 

import (
	"errors"
	"fmt"
	"context"
)

func commandMapf(cfg *config) error {
  locations,err,statusCode := cfg.pokeapiClient.ListLocations(cfg.NextLocationsURL,context.Background())
  if err != nil && statusCode == 404 {
	return fmt.Errorf("oops! failed to get the areas")
  } else if err != nil {
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

	locations,err,statusCode := cfg.pokeapiClient.ListLocations(cfg.PrevLocationsURL,context.Background())
	if err != nil && statusCode == 404{
		return fmt.Errorf("oops! can't get the areas")
	} else if err != nil {
		return err
	}

	cfg.NextLocationsURL = locations.Next 
	cfg.PrevLocationsURL = locations.Previous

	for _,loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}