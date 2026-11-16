package main 

import (
	"fmt"
)

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex! \n")
	fmt.Printf("Usage:\n")

	fmt.Printf("\n")

	fmt.Printf("help: Displays a help message \n")
	fmt.Printf("exit: Exit the Pokedex")
	return nil
}