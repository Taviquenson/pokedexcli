package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Taviquenson/pokedexcli/internal/pokecache"
)

var exploreCache = pokecache.NewCache(5 * time.Second)

func Explore(config *Config, cmdParams ...string) error {
	if len(cmdParams) == 0 {
		return fmt.Errorf("error: Command requires an area name.\n  e.g.\n\texplore pastoria-city-area")
	}
	areaName := cmdParams[0]
	bodyEntry, exists := exploreCache.Get(areaName)
	if exists { // don't add to cache
		listPokemon(bodyEntry)
		// fmt.Println("Was in cache")
		return nil
	} else { // add to cache
		res, err := http.Get(BaseURL + "/" + areaName)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode == 404 {
			return fmt.Errorf("error: Area name not found.\nFor a list of valid areas use command 'map'")
		} else if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		exploreCache.Add(areaName, body)
		listPokemon(body)
		// fmt.Println("Wasn't in cache")
		return nil
	}
}

func listPokemon(body []byte) error {
	location := Location{}
	err := json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", location.Name)
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
