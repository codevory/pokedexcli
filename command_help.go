package main 

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex! \n")
	fmt.Printf("Usage:\n")

	fmt.Printf("\n")

	for k,v := range cfg.cliCommandMap {
		fmt.Printf("%s : %s\n",k,v.description)
	}

	return nil
}