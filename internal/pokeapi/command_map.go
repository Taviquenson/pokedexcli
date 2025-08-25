package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PokeMap(config *Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
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
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	return nil
}
