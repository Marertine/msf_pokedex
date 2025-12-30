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
		outputString := fmt.Sprintf("Your command was: %s\n", firstWord)
		fmt.Print(outputString)

	}
}

func cleanInput(text string) []string {
	var sliceStrings []string

	loweredString := strings.ToLower(text)
	sliceStrings = strings.Fields(loweredString)

	return sliceStrings
}
