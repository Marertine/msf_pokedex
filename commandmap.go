package main

import (
	"fmt"
	//"msf_pokedex/internal/pokeapi"
	"os"
	"time"

	"github.com/Marertine/msf_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeConfig) error
}

type pokeConfig struct {
	nextURL string
	prevURL string
}

var pokeClient = pokeapi.NewClient(5 * time.Second)

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
			description: "Display 20 locations at a time",
			callback:    commandMap,
		},
	}
}

func commandExit(cfg *pokeConfig) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeConfig) error {
	mapCommands := getCommands()
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, commandItem := range mapCommands {
		fmt.Printf("%s: %s\n", commandItem.name, commandItem.description)
	}

	return nil
}

func commandMap(cfg *pokeConfig) error {
	fmt.Print("FFS\n")
	return nil
}
