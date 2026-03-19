package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	cfgPtr := &cfg
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())

		if len(text) == 0 {
			continue
		}

		val, ok := getCommands()[text[0]]
		if ok {
			err := val.callback(cfgPtr)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
}

type Response struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "shows the next 20 poke place maps (mapb to go back)",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "shows the previous 20 poke place maps (mapb to go back)",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
