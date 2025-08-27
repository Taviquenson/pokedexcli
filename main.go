package main

import (
	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
)

func main() {
	config := pokeapi.PokeConfig
	startRepl(&config)
}
