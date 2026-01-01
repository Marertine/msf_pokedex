package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/Marertine/msf_pokedex/internal/pokeapi"
	//"time"
	//"github.com/Marertine/msf_pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type pokeConfig struct {
	pokeapiClient pokeapi.Client
	nextURL       string
	prevURL       string
	caughtPokemon map[string]pokeapi.Pokemon
}

//var pokeClient = pokeapi.NewClient(5 * time.Second)
//var caughtPokemon map[string]pokeapi.Pokemon

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Catch a monster",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Show list of Pokemon in a location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "See details about a Pokemon",
			callback:    commandInspect,
		}, "map": {
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

func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	// _ indicates that we are intentionally not using these parameters
	mapCommands := getCommands()
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, commandItem := range mapCommands {
		fmt.Printf("%s: %s\n", commandItem.name, commandItem.description)
	}

	return nil
}

func commandMap(cfg *config, args ...string) error {
	var pageURL *string
	if cfg.nextURL != "" {
		pageURL = &cfg.nextURL
	}

	resp, err := cfg.pokeapiClient.ListLocations(pageURL)
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

func commandMapB(cfg *config, args ...string) error {
	var pageURL *string
	if cfg.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	pageURL = &cfg.prevURL
	resp, err := cfg.pokeapiClient.ListLocations(pageURL)
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

func commandExplore(cfg *config, args ...string) error {
	//fmt.Println("DEBUG: explore called with args:", words)

	if len(args) != 1 {
		//fmt.Println("DEBUG: wrong number of args")
		return errors.New("you must provide a location name")
	}

	name := args[0]
	//fmt.Println("DEBUG: fetching location:", name)

	location, err := cfg.pokeapiClient.GetLocation(name)
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

func commandCatch(cfg *config, args ...string) error {
	//fmt.Println("DEBUG: catch called with args:", words)

	if len(args) != 1 {
		//fmt.Println("DEBUG: wrong number of args")
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	//fmt.Println("DEBUG: catching pokemon:", name)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		//fmt.Println("DEBUG: GetPokemon error:", err)
		return err
	}

	//fmt.Println("DEBUG: monster name:", pokemon.Name)
	//fmt.Println("DEBUG: base experience:", pokemon.BaseExperience)

	intCatchChance := max(min(80-(pokemon.BaseExperience/10), 95), 10)

	if rand.Intn(100) < intCatchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
