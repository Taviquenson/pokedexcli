package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() // advances the Scanner to the next line

		userWords := cleanInput(scanner.Text())
		if len(userWords) == 0 { // Don't print message if scanned line was empty
			continue
		}

		command := userWords[0]

		fmt.Printf("Your command was: %v\n", command)
	}
}

func cleanInput(text string) []string {
	subStrs := strings.Fields(text)
	for i, str := range subStrs {
		subStrs[i] = strings.ToLower(str)
	}
	return subStrs
}
