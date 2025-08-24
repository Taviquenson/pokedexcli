package main

import "fmt"

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
