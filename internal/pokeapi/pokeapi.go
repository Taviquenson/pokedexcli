package pokeapi

type Config struct {
	Next     string
	Previous *string
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
