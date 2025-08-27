package main

import (
	// "fmt"
	"github.com/Taviquenson/pokedexcli/internal/pokeapi"
	// "internal/pokecache"
	// "time"
)

func main() {
	config := pokeapi.PokeConfig
	// pokeCache := pokecache.NewCache(5 * time.Second)
	// pokeCache.Add("https://example.com", []byte("testdata"))
	// val, wasFound := pokeCache.Get("https://example.com")
	// fmt.Printf("\nValue: %v and bool: %v\n", string(val), wasFound)
	// go func() {
	// 	time.Sleep(7 * time.Second)
	// 	val, wasFound = pokeCache.Get("https://example.com")
	// 	fmt.Printf("\nValue: %v and bool: %v\n", string(val), wasFound)
	// }()
	startRepl(&config)
}
