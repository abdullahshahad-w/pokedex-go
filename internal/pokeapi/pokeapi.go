package pokeapi

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"io"
	"github.com/abdullahshahad-w/pokedex-go/internal/pokecache"
)

type CliCommand struct {
	Name string
	Description string
	Callback func(*Config) error
}

type Config struct {
	CommandRegistry map[string]CliCommand
	NextUrl string
	PreviousUrl string
	Cache *pokecache.Cache
}

type locationAreaRes struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
}

func CommandExit(reg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(reg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for key, value := range reg.CommandRegistry {
		fmt.Printf("%s: %s\n", key, value.Description)
	}

	return nil
}

func CommandMap(reg *Config) error {
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

func CommandMapb(reg *Config) error {
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