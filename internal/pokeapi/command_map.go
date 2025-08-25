package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PokeMap(config *Config) error {
	// var res any
	// var err any
	res, err := mapRequest(config)
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

	// firstMaps := "https://pokeapi.co/api/v2/location-area"
	// firstMaps := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

	// fmt.Printf("config's Previous: %s\n", *config.Previous)
	// if locationAreas.Previous == nil {
	// 	config.Previous = &firstMaps
	// } else {
	// 	// fmt.Println("NEVER HERE")
	// 	config.Previous = locationAreas.Previous
	// }
	// fmt.Printf("config's new Previous: %s\n", *config.Previous)

	// fmt.Printf("config's Next: %s\n", config.Next)
	config.Next = locationAreas.Next
	fmt.Printf("config's new Next: %s\n", config.Next)

	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}

	return nil
}

func mapRequest(config *Config) (*http.Response, error) {
	// firstMaps := "https://pokeapi.co/api/v2/location-area"
	firstMaps := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	// currentMaps :=
	var res *http.Response
	var err error
	if config.Previous == nil {
		// fmt.Printf("config's Previous: %v\n", config.Previous)
		// fmt.Printf("config's Next: %v\n", config.Next)
		config.Previous = &firstMaps
		fmt.Printf("config's new Previous: %s\n", *config.Previous)
		res, err = http.Get(*config.Previous)
	} else {
		// fmt.Printf("in mapRequest(), config.Next: %s\n", config.Next)
		// fmt.Printf("config's Previous: %s\n", *config.Previous)
		// fmt.Printf("config's Next: %s\n", config.Next)
		config.Previous = &config.Next
		fmt.Printf("config's new Previous: %s\n", *config.Previous)
		res, err = http.Get(config.Next)
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return res, err
}
