package pokedex

import (
	"fmt"

	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
)

func ListPokedex(config *pokeapi.Config, cmdParams ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokeapi.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
