# Pokedex Go

A lightweight terminal-based Pokédex built in Go. This project lets you explore Pokémon locations from the public PokeAPI, catch creatures, inspect their details, and manage a personal collection in a simple REPL-style command prompt.

## Features

- Interactive Pokédex CLI
- Browse Pokémon regions and locations
- Explore a specific area for Pokémon encounters
- Catch wild Pokémon with a chance-based mechanic
- Inspect a caught Pokémon's stats and types
- View your collection using the `pokedex` command
- Basic HTTP response caching to reduce repeated API calls

## Project Structure

```text
.
├── main.go                 # Application bootstrap and command registration
├── repl.go                 # Read-eval-print loop for user interaction
├── repl_test.go            # Tests for command input sanitization
├── go.mod                  # Go module definition
├── internal/
│   ├── pokeapi/
│   │   └── pokeapi.go      # API client, commands, and Pokémon logic
│   └── pokecache/
│       └── cache.go        # In-memory cache for API responses
```

## Commands

Once the application is running, the following commands are available:

- `help` — Display all available commands
- `map` — Show the next page of Pokémon locations
- `mapb` — Show the previous page of Pokémon locations
- `explore <location>` — List Pokémon in a specific area
- `catch <pokemon>` — Attempt to catch a Pokémon
- `inspect <pokemon>` — Show details for a caught Pokémon
- `pokedex` — Show all Pokémon caught so far
- `exit` — Exit the program

## Getting Started

### Prerequisites

- Go 1.27.1 or a compatible newer version

### Run the app

```bash
go run .
```

### Example session

```text
Pokedex > help
Welcome to the Pokédex!
Usage:
exit: Exits the Pokedex
help: Shows how to use the Pokedex
map: Shows the locations of the Pokemon world
mapb: Shows the previous locations of the Pokemon world
explore: Shows pokemons in the area
catch: Tries to catch a pokemon
inspect: Shows info of a pokemon
pokedex: Shows all the pokemon you've caught

Pokedex > map
pallet-town
viridian-city
...

Pokedex > explore pallet-town
Found Pokemon:
 - bulbasaur
 - squirtle

Pokedex > catch bulbasaur
Throwing a Pokeball at bulbasaur...

bulbasaur was caught!

Pokedex > inspect bulbasaur
Name: bulbasaur
Height: 7
Weight: 69
Stats:
 -hp: 45
 -attack: 49
...
Types:
 - grass
 - poison

Pokedex > pokedex
Your Pokedex:
 - bulbasaur
```

## Notes

This project uses the public PokeAPI and demonstrates core Go concepts such as:

- HTTP requests
- JSON decoding
- concurrency with a cache reaper loop
- CLI interactive programming
- basic testing

## License

This project does not currently include a license file. If you plan to share or distribute it publicly, consider adding an open-source license such as MIT.

## Acknowledgements

- [PokeAPI](https://pokeapi.co/) for the Pokémon data
- Go standard library for CLI and HTTP functionality
