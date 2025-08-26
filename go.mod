module github.com/Taviquenson/pokedexcli

go 1.24.4

require internal/pokeapi v1.24.4
replace internal/pokeapi => ./internal/pokeapi
require internal/pokecache v1.24.4
replace internal/pokecache => ./internal/pokecache