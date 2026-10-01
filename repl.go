package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/abdullahshahad-w/pokedex-go/internal/pokeapi"
)

func cleanInput(text string) []string {
	result := []string{}
	if len(text) == 0 {
		return result
	}
	
	words := strings.Fields(text)
	

	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}

	return result
}

func replLoop(reg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for ;; {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		command := cleanedInput[0]

		if val, ok := reg.CommandRegistry[command]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := val.Callback(reg)
			if err != nil {
				fmt.Printf("error calling callback function for command: %s, error: %v", command, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading the input: %v\n", err)
	}
}