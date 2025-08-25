package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PokeMapB(config *Config) error {
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
	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println(err)
	}

	config.Previous = locationAreas.Previous
	config.Next = locationAreas.Next
	// if locationAreas.Previous == nil {
	// 	fmt.Printf("config's new Previous: %v\n", config.Previous)
	// } else {
	// 	fmt.Printf("config's new Previous: %s\n", *config.Previous)
	// }
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func PokeMap(config *Config) error {
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

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println(err)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	// fmt.Printf("config's new Next: %s\n", config.Next)
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapRequest(config *Config, isMapb bool) (*http.Response, error) {
	var res *http.Response
	var err error
	if config.Previous == nil {
		if isMapb {
			return res, err
		} else {
			// fmt.Printf("@First Requesting Next: %v\n", config.Next)
			res, err = http.Get(config.Next)
			return res, err
		}
	} else {
		if isMapb {
			// fmt.Printf("@Other Requesting Previous: %s\n", *config.Previous)
			res, err = http.Get(*config.Previous)
			return res, err
		} else {
			// fmt.Printf("@Other Requesting Next: %v\n", config.Next)
			res, err = http.Get(config.Next)
			return res, err
		}
	}
}
