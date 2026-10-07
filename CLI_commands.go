package main

import (
	"errors"
	"fmt"
	"math/rand"
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

func commandExit(cfg *config, _ []string) error {
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
	fmt.Println()

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

func catch(cfg *config, Poke []string) error {
	item, err := cfg.pokeapiClient.PokemonDetails(Poke[1])
	if err != nil {
		return err
	}
	roll := rand.Intn(item.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...\n", item.Name)
	if roll > 50 {
		fmt.Printf("%v escaped\n", item.Name)

	} else {
		fmt.Printf("%v was caught!\n", item.Name)
		cfg.pokedex[item.Name] = item
	}

	return nil
}

func inspect(cfg *config, Poke []string) error {
	if len(Poke) < 2 {
		fmt.Printf("Need to name a pokemon!")
		return nil
	}
	if pokemon, ok := cfg.pokedex[Poke[1]]; !ok {
		fmt.Printf("you have not caught that pokemon")
		return nil
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	}

	return nil
}
func pokedexcli(cfg *config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil

}
