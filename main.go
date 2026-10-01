package main

func main() {
	config := config{
		commandRegistry: map[string]cliCommand{
			"exit": {
				name: "exit",
				description: "Exits the Pokedex",
				callback: commandExit,
			},
			"help": {
				name: "help",
				description: "Shows how to use the Pokedex",
				callback: commandHelp,
			},
		},
	}

	replLoop(&config)
}