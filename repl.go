
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codevory/pokedexcli/pokeapi"
)

func cleanInput(text string) []string {
	var res []string
	if len(text) < 1 {
		return res
	}

   for _,val := range strings.Fields(text) {
	var loweredCase string = strings.ToLower(val)
	res = append(res,loweredCase)
   }
	return res
}

type cliCommand struct {
	name string 
	description string 
	callback func(config *config) error
}


type config struct{
	cliCommandMap map[string]cliCommand
	pokeapiClient pokeapi.Client
	NextLocationsURL *string 
	PrevLocationsURL *string 
}

func startRepl(cfg *config){

  scanner := bufio.NewScanner(os.Stdin)

  for {
    fmt.Print("Pokedex > ")
	exist := scanner.Scan()

	if !exist {
	  fmt.Errorf("unexpected error getting user input")
	}

	text := scanner.Text()

	words := cleanInput(text)
	for _,keyword := range words {
			if _,ok := cfg.cliCommandMap[keyword]; ok {
				if err := cfg.cliCommandMap[keyword].callback(cfg); err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Print("Unknown command")
			}
	}
	fmt.Printf("\n")
  }
}


func getCommands()map[string]cliCommand {
	return map[string]cliCommand{
      "help": {
	    	name: "help",
	    	description: "get info about pokedexcli",
	    	callback: commandHelp,
	    },
		"map": {
			name: "map",
			description: "gets list of 20 pokemon locations",
			callback: commandMapf,
		},
		"mapb": {
			name: "map back",
			description: "gets list of 20 previous locations",
			callback: commandMapB,
		},
		"exit": {
		  name: "exit",
		  description: "Exit the Pokedex",
		  callback: commandExit,
	    },
	}
}