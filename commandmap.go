package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Marertine/msf_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeConfig, []string) error
}

type pokeConfig struct {
	nextURL string
	prevURL string
}

var pokeClient = pokeapi.NewClient(5 * time.Second)

func getCommands(words []string) map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Show list of Pokemon in a location",
			callback:    commandExplore,
		},
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

func commandExit(_ *pokeConfig, _ []string) error {
	// _ indicates that we are intentionally not using these parameters
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(_ *pokeConfig, words []string) error {
	// _ indicates that we are intentionally not using these parameters
	mapCommands := getCommands(words)
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, commandItem := range mapCommands {
		fmt.Printf("%s: %s\n", commandItem.name, commandItem.description)
	}

	return nil
}

func commandMap(cfg *pokeConfig, _ []string) error {
	// _ indicates that we are intentionally not using these parameters
	var pageURL *string
	if cfg.nextURL != "" {
		pageURL = &cfg.nextURL
	}

	resp, err := pokeClient.ListLocations(pageURL)
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

func commandMapB(cfg *pokeConfig, _ []string) error {
	// _ indicates that we are intentionally not using these parameters
	var pageURL *string
	if cfg.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	pageURL = &cfg.prevURL
	resp, err := pokeClient.ListLocations(pageURL)
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

func commandExplore(cfg *pokeConfig, words []string) error {
	//fmt.Println("DEBUG: explore called with args:", words)

	if len(words) < 2 {
		//fmt.Println("DEBUG: wrong number of args")
		return errors.New("you must provide a location name")
	}

	// command is words[0], location name is words[1]
	name := words[1]
	//fmt.Println("DEBUG: fetching location:", name)

	location, err := pokeClient.GetLocation(name)
	if err != nil {
		//fmt.Println("DEBUG: GetLocation error:", err)
		return err
	}

	//fmt.Println("DEBUG: location name:", location.Name)
	//fmt.Println("DEBUG: num encounters:", len(location.PokemonEncounters))

	//fmt.Printf("Exploring %s...\n", location.Name)
	//fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
