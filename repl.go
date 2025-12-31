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

func startRepl() {
	cfg := &pokeConfig{}
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
		if cmd, found := mapCommands[firstWord]; found {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
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
