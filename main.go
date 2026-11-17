package main 

import (
	"time"

	"github.com/codevory/pokedexcli/pokeapi"
)

func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		cliCommandMap: getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
