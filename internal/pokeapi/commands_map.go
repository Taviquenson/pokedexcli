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

var pokeCache = pokecache.NewCache(5 * time.Second)

func Maps(config *Config) error {
	bodyEntry, exists := pokeCache.Get(config.Next)
	if exists {
		fmt.Println("Not saved to cache")
		listLocations(bodyEntry, config)
		return nil
	} else { // add to cache
		res, err := mapRequest(config, false)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Saved to cache")
		pokeCache.Add(config.Next, body)
		listLocations(body, config)
		return nil
	}

}

func MapsB(config *Config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	bodyEntry, exists := pokeCache.Get(*config.Previous)
	if exists {
		fmt.Println("Not saved to cache")
		listLocations(bodyEntry, config)
		return nil
	} else { // add to cache
		res, err := mapRequest(config, true)
		if err != nil {
			log.Fatal(err)
		}
		if res == nil {
			fmt.Println("you're on the first page")
			return nil
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		pokeCache.Add(*config.Previous, body)
		fmt.Println("Saved to cache")
		listLocations(body, config)
		return nil
	}
}

// Makes the adequate request whther advancing or bactracking in the maps
func mapRequest(config *Config, isMapb bool) (*http.Response, error) {
	var res *http.Response
	var err error
	if config.Previous == nil {
		if isMapb {
			return res, err
		} else {
			res, err = http.Get(config.Next)
			return res, err
		}
	} else {
		if isMapb {
			res, err = http.Get(*config.Previous)
			return res, err
		} else {
			res, err = http.Get(config.Next)
			return res, err
		}
	}
}

// Also updates the Config URLs
func listLocations(body []byte, config *Config) error {
	locationAreas := LocationAreas{}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println(err)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	return nil
}
