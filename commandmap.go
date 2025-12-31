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
		"mapb": {
			name:        "mapb",
			description: "Go to previous page of 20 locations",
			callback:    commandMapB,
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
	var pageURL *string
	if cfg.nextURL != "" {
		pageURL = &cfg.nextURL
	}

	resp, err := pokeClient.ListLocations(pageURL)
	//resp, err := pokeClient.ListLocations(nil)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	if resp.Next != nil {
		cfg.nextURL = *resp.Next
	} else {
		cfg.nextURL = ""
	}

	if resp.Previous != nil {
		cfg.prevURL = *resp.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}

func commandMapB(cfg *pokeConfig) error {
	var pageURL *string
	if cfg.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	pageURL = &cfg.prevURL
	resp, err := pokeClient.ListLocations(pageURL)
	//resp, err := pokeClient.ListLocations(nil)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	if resp.Next != nil {
		cfg.nextURL = *resp.Next
	} else {
		cfg.nextURL = ""
	}

	if resp.Previous != nil {
		cfg.prevURL = *resp.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}
