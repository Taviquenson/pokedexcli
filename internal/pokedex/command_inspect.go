package pokedex

import (
	"fmt"

	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
)

func Inspect(config *pokeapi.Config, cmdParams ...string) error {
	if len(cmdParams) == 0 {
		return fmt.Errorf("error: Command requires a pokemon name.\n  e.g.\n\tinspect pikachu")
	}
	pokemonName := cmdParams[0]
	pokemon, exists := pokeapi.Pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	listPokeInfo(pokemon)

	return nil
}

func listPokeInfo(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  -%s\n", pType.Type.Name)
	}
}
