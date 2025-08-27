package main

import (
	"bufio"
	"fmt"
	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

func startRepl(config *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() // advances the Scanner to the next line

		userWords := cleanInput(scanner.Text())
		if len(userWords) == 0 { // Don't print message if scanned line was empty
			continue
		}

		commandName := userWords[0]

		command, ok := getCommands()[commandName]
		if ok {
			// var config *pokeapi.Config // this is not initialized
			err := command.callback(config)
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

func cleanInput(text string) []string {
	subStrs := strings.Fields(text)
	for i, str := range subStrs {
		subStrs[i] = strings.ToLower(str)
	}
	return subStrs
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

// Using a function that returns the map of commands to avoid having
// to deal with an initialization loop that would happen because
// in the map there's a reference to the help command and in the help
// command there would be a reference to the map
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows a list of locations and advances the list 20 at a time",
			callback:    pokeapi.Maps,
		},
		"mapb": {
			name:        "mapb",
			description: "Moves back in a list of locations 20 at a time",
			callback:    pokeapi.MapsB,
		},
	}
}
