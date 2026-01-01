package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Marertine/msf_pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       string
	prevURL       string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		commands := getCommands()
		cmd, found := commands[commandName]
		if !found {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg, args...); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	var sliceStrings []string

	loweredString := strings.ToLower(text)
	sliceStrings = strings.Fields(loweredString)

	return sliceStrings
}
