package main

import (
	"fmt"

	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config, cmdParams ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
