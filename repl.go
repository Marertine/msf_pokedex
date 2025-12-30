package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"io"
	//"net/http"
	//"bytes"
	//"encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		userInput := scanner.Text()

		words := cleanInput(userInput)

		if len(words) == 0 {
			continue
		}

		firstWord := words[0]

		mapCommands := getCommands()
		if value, found := mapCommands[firstWord]; found {
			value.callback()
		} else {
			//outputString := fmt.Sprintf("Your text was: %s\n", firstWord)
			fmt.Print("Unknown command\n")
		}

	}
}

func cleanInput(text string) []string {
	var sliceStrings []string

	loweredString := strings.ToLower(text)
	sliceStrings = strings.Fields(loweredString)

	return sliceStrings
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	mapCommands := getCommands()
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, commandItem := range mapCommands {
		fmt.Printf("%s: %s\n", commandItem.name, commandItem.description)
	}

	return nil
}
