package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) PokemonDetails(name string) (Pokemon, error) {
	var requestURL string
	if name == "" {
		return Pokemon{}, errors.New("Invalid pokemon")
	} else {
		requestURL = "https://pokeapi.co/api/v2/pokemon/" + name
	}
	if val, ok := c.cache.Get(requestURL); ok {
		var cached Pokemon
		err := json.Unmarshal(val, &cached)
		if err != nil {
			return Pokemon{}, err
		}
		return cached, nil
	}

	resp, err := http.Get(requestURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(requestURL, body)
	return pokemon, nil
}
