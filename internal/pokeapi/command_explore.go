package pokeapi

import (
	"github.com/Taviquenson/pokedexcli/internal/pokecache"
	"time"
)

var exploreCache = pokecache.NewCache(5 * time.Second)

func Explore(config *Config, cmdParams ...string) error {
	exploreCache.Add("kupo", []byte("kupo kupo!"))
	return nil
}
