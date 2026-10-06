package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"errors"
)

type pokemons struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
func (c *Client) Getpokemons(name string) (pokemons, error) {
	var requestURL string
	if name == "" {
		return pokemons{},errors.New("Invalid location")
	} else {
		requestURL = "https://pokeapi.co/api/v2/location-area/"+name
	}
	if val, ok := c.cache.Get(requestURL); ok {
		var cached pokemons
		err := json.Unmarshal(val, &cached)
		if err != nil {
			return pokemons{}, err
		}
		return cached, nil
	}

	resp, err := http.Get(requestURL)
	if err != nil {
		return pokemons{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemons{}, err
	}

	var pokemon pokemons
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return pokemons{}, err
	}
	c.cache.Add(requestURL, body)
	return pokemon, nil
}