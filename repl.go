package main

import (
	"github.com/dariusss185/Pokedex/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	commands      map[string]cliCommand
	next          *string
	previous      *string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
