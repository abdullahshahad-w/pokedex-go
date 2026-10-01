package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	commandRegistry map[string]cliCommand
}

func commandExit(reg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(reg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for key, value := range reg.commandRegistry {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}