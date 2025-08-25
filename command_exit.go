package main

import (
	"fmt"
	"internal/pokeapi"
	"os"
)

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
