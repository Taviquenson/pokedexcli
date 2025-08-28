package pokeapi

import (
	"encoding/json"
	"fmt"
)

var pokedex = make(map[string]Pokemon)

func Catch(config *Config, cmdParams ...string) error {
	if len(cmdParams) == 0 {
		return fmt.Errorf("error: Command requires a pokemon name.\n  e.g.\n\tcatch pikachu")
	}
	pokemonName := cmdParams[0]
	body, err := getBody("pokemon", pokemonName)
	if err != nil {
		return err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println(err)
	}

	// pokemon with highest base experience is Blissey @ 608
	// pokemon with lowest base experience is Sunkern @ 36

	pokedex[pokemonName] = pokemon

	fmt.Println(pokedex[pokemonName])

	return nil
}
