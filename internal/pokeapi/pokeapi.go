package pokeapi

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/abdullahshahad-w/pokedex-go/internal/pokecache"
)

type CliCommand struct {
	Name string
	Description string
	Callback func(*Config, string) error
}

type Config struct {
	CommandRegistry map[string]CliCommand
	NextUrl string
	PreviousUrl string
	Cache *pokecache.Cache
	Pokedex map[string]PokeCatch
}

type locationAreaRes struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
}

type pokemonEncounter struct {
	Pokemons []pokemonEntry `json:"pokemon_encounters"`
}

type pokemonEntry struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type PokeCatch struct {
	Name string `json:"name"`
	Exp int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		}
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func CommandExit(reg *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(reg *Config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for key, value := range reg.CommandRegistry {
		fmt.Printf("%s: %s\n", key, value.Description)
	}

	return nil
}

func CommandMap(reg *Config, arg string) error {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if reg.NextUrl != "" {
		url = reg.NextUrl
	}

	var locationRes locationAreaRes	

	if raw, ok := reg.Cache.Get(url); ok {
		if err := json.Unmarshal(raw, &locationRes); err != nil {
			return err
		}

	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		reg.Cache.Add(url, data)


		if err := json.Unmarshal(data, &locationRes); err != nil {
			return err
		}
	}

		reg.NextUrl = locationRes.Next
		reg.PreviousUrl = locationRes.Previous

		for _, location := range locationRes.Results {
			fmt.Println(location.Name)
		}

		return nil
}

func CommandMapb(reg *Config, arg string) error {
	if reg.PreviousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := reg.PreviousUrl

	var locationRes locationAreaRes

	if raw, ok := reg.Cache.Get(url); ok {
		if err := json.Unmarshal(raw, &locationRes); err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		reg.Cache.Add(url, data)

		if err := json.Unmarshal(data, &locationRes); err != nil {
			return err
		}
	}

	reg.NextUrl = locationRes.Next
	reg.PreviousUrl = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandExplore(reg *Config, arg string) error {
	fmt.Printf("\nExploring %s...\n", arg)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", arg)

	var pokemonEncounter pokemonEncounter

	if raw, ok := reg.Cache.Get(url); ok {
		if err := json.Unmarshal(raw, &pokemonEncounter); err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		reg.Cache.Add(url, data)

		if err := json.Unmarshal(data, &pokemonEncounter); err != nil {
			return err
		}
	}

	fmt.Println("Found Pokemon:")

	for _, entry := range pokemonEncounter.Pokemons {
		fmt.Printf(" - %s\n", entry.Pokemon.Name)
	}

	return nil
}

func CommandCatch(reg *Config, arg string) error {
	if _, ok := reg.Pokedex[arg]; ok {
		fmt.Println("You have already caught the pokemon!")
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", arg)

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var catch PokeCatch

	if err := json.Unmarshal(data, &catch); err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", arg)

	rand := rand.Float64()
	difficulty := 0.01
	prob := 1 / (1 + difficulty * float64(catch.Exp))

	if rand < prob {
		reg.Pokedex[arg] = catch
		fmt.Printf("\n%s was caught!\n", arg)
	} else {
		fmt.Printf("\n%s escaped!\n", arg)
	}

	return nil
}

func CommandInspect(reg *Config, arg string) error {
	if _, ok := reg.Pokedex[arg]; !ok {
		fmt.Println("You have not caught this pokemon!")
		return nil
	}

	pokemon := reg.Pokedex[arg]

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)		
	}

	return nil
}

func CommandPokedex(reg *Config, arg string) error {
	if len(reg.Pokedex) == 0 {
		fmt.Println("You have not yet caught any pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key := range reg.Pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}