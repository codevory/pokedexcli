package main

import (
	"bufio"
	"fmt"
	"github.com/codevory/pokedexcli/pokeapi"
	"github.com/codevory/pokedexcli/pokecache"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var res []string
	if len(text) < 1 {
		return res
	}

	for _, val := range strings.Fields(text) {
		var loweredCase string = strings.ToLower(val)
		res = append(res, loweredCase)
	}
	return res
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	cliCommandMap    map[string]cliCommand
	pokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	areaName         string
	cache            *pokecache.Cache
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		exist := scanner.Scan()

		if !exist {
			fmt.Errorf("unexpected error getting user input")
		}

		text := scanner.Text()

		words := cleanInput(text)
		var keyword string = words[0]
		if len(words) >= 2 {
			cfg.areaName = words[1]
		}

		if _, ok := cfg.cliCommandMap[keyword]; ok {
			if err := cfg.cliCommandMap[keyword].callback(cfg); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Print("Unknown command \n")
		}
	}
	fmt.Printf("\n")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "get info about pokedexcli",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "gets list of 20 pokemon locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map back",
			description: "gets list of 20 previous locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explore the available pokemons available in area",
			callback:    explore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
