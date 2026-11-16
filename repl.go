
package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
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
	callback func() error
}


func startRepl(){

	cliCommandMap := map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "get info about pokedexcli",
			callback: commandHelp,
		},
	}

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
			if _,ok := cliCommandMap[keyword]; ok {
				cliCommandMap[keyword].callback()
			} else {
				fmt.Print("Unknown command")
			}
	}
	fmt.Printf("\n")
  }
}

func main(){
	startRepl()
}
