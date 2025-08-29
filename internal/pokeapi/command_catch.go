package pokeapi

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Create a local RNG with a dynamic seed
// (to control seeding for reproducibility or isolated RNG streams)
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

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

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if wasCaught(pokemon.BaseExperience) {
		Pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", Pokedex[pokemonName].Name)
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// Catches a pokemon based on its base experience attribute
// The pokemon with the highest base experience is Blissey at 608
// The pokemon with the lowest base experience is Sunkern at 36
func wasCaught(exp int) bool {
	// Clamp exp to the range [36, 608]
	minX, maxX := 36.0, 608.0
	clamped := math.Min(math.Max(float64(exp), minX), maxX)

	// Normalize exp into a range [0, 1]
	// t=0 means near 36, t=1 means near 608
	t := (clamped - minX) / (maxX - minX)

	// Introduce bias
	bias := math.Pow(rng.Float64(), t)

	// Map bias [0,1] to [0.25, 1.0]
	minVal, maxVal := 0.25, 1.0
	captureChance := minVal + (maxVal-minVal)*bias
	// fmt.Println(captureChance)
	return captureChance > 0.75
}
