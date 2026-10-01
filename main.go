package main

import (
	"time"
	"github.com/abdullahshahad-w/pokedex-go/internal/pokeapi"
	"github.com/abdullahshahad-w/pokedex-go/internal/pokecache"
)
func main() {
	cache := pokecache.NewCache(30 * time.Second)
	config := pokeapi.Config{
		CommandRegistry: map[string]pokeapi.CliCommand{
			"exit": {
				Name: "exit",
				Description: "Exits the Pokedex",
				Callback: pokeapi.CommandExit,
			},
			"help": {
				Name: "help",
				Description: "Shows how to use the Pokedex",
				Callback: pokeapi.CommandHelp,
			},
			"map": {
				Name: "map",
				Description: "Shows the locations of the Pokemon world",
				Callback: pokeapi.CommandMap,
			},
			"mapb": {
				Name: "mapb",
				Description: "Shows the previous locations of the Pokemon world",
				Callback: pokeapi.CommandMapb,
			},
		},
		Cache: cache,
	}

	replLoop(&config)
}