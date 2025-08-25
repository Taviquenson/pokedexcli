package pokeapi

type Config struct {
	Next     string
	Previous *string
}

var PokeConfig = Config{
	Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous: nil,
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"` // using *string 'cuz it can be null value
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
