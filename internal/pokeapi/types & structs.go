package pokeapi

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type LocationAreaResponse struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

