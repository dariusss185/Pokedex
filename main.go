package main

import (
	"bufio"
	"fmt"
	"github.com/dariusss185/Pokedex/internal/pokeapi"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)

	cfg := config{
		pokedex:       make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeClient,
		commands: map[string]cliCommand{
			"inspect": cliCommand{
				name:        "inspect",
				description: "Shows the name, height, weight, stats of the Pokemon",
				callback:    inspect,
			},
			"exit": cliCommand{
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "print the usage menu",
				callback:    initialize,
			},
			"map": {
				name:        "map",
				description: "display a list of next 20 locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "display a list of 20 previous locations",
				callback:    commandMapB,
			},
			"explore": {
				name:        "explore",
				description: "list pokemons inside a location",
				callback:    explore,
			},
			"catch": {
				name:        "catch",
				description: "catch a pokemon!",
				callback:    catch,
			},
			"pokedex": {
				name:        "pokedex",
				description: "list all the pokemons you have caught!",
				callback:    pokedexcli,
			},
		},
	}
	initialize(&cfg, nil)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) > 0 {
			if command, exists := cfg.commands[words[0]]; exists {
				err := command.callback(&cfg, words)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
		continue
	}
}
