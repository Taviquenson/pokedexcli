package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Maps(config *Config) error {
	res, err := mapRequest(config, false)
	if err != nil {
		log.Fatal(err)
	}

	listLocations(res, config)

	return nil
}

func MapsB(config *Config) error {
	res, err := mapRequest(config, true)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	listLocations(res, config)

	return nil
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
func listLocations(res *http.Response, config *Config) error {
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
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
