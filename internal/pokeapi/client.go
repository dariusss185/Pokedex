package pokeapi

import (
	"encoding/json"
	"github.com/dariusss185/Pokedex/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
	}
}

func (c *Client) GetLocationAreas(url *string) (LocationAreaResponse, error) {
	var requestURL string

	if url == nil {
		requestURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		requestURL = *url
	}
	if val, ok := c.cache.Get(requestURL); ok {
		var cached LocationAreaResponse
		err := json.Unmarshal(val, &cached)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return cached, nil
	}

	resp, err := http.Get(requestURL)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreas LocationAreaResponse
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(requestURL, body)
	return locationAreas, nil
}
