package main

import (
	"errors"
	"fmt"
	"os"
)

func commandMap(cfg *config, _ []string) error {
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cfg *config, _ []string) error {
	if cfg.previous == nil {
		return errors.New("You are on the first page")
	}
	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExit(cfg *config,_ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func initialize(cfg *config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func explore(cfg *config, location []string) error {
	Pokemons, err := cfg.pokeapiClient.Getpokemons(location[1])
	if err != nil {
		return err
	}
	for _, data := range Pokemons.Encounters {
		fmt.Println(data.Pokemon.Name)
	}
	return nil
}
