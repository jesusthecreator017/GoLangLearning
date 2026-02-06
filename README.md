# GoLangLearning

A Go learning repository featuring a Pokedex CLI application and backend HTTP learning modules.

## Pokedex CLI

An interactive command-line Pokedex that uses the [PokeAPI](https://pokeapi.co/) to explore locations, catch Pokemon, and manage your collection.

### Features

- **Explore the Pokemon world** - Navigate through location areas with paginated results
- **Catch Pokemon** - Attempt to catch Pokemon with probability based on their base experience
- **Build your Pokedex** - View stats, types, and abilities of your caught Pokemon
- **Caching** - Built-in response caching to minimize API calls

### Quick Start

```bash
cd pokedex
go build
./pokedex
```

### Commands

| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `map` | Show next 20 location areas |
| `mapb` | Show previous 20 location areas |
| `explore <location>` | List Pokemon in a location area |
| `catch <pokemon>` | Attempt to catch a Pokemon |
| `inspect <pokemon>` | View details of a caught Pokemon |
| `pokedex` | List all caught Pokemon |
| `exit` | Exit the application |

### Example Session

```
Pokedex > map
 - canalave-city-area
 - eterna-city-area
 - pastoria-city-area
 ...

Pokedex > explore pastoria-city-area
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 ...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  - hp: 20
  - attack: 10
  ...
Types:
  - water
```

## Backend Learning Modules

The `backend/` directory contains learning exercises covering:

- **DNS** - URL sections, web addresses, protocol handling
- **HTTP Methods** - GET, POST, PUT, DELETE implementations
- **JSON** - Syntax, decoding, marshaling, mapped interfaces
- **Headers** - HTTP headers and their uses
- **Paths** - URL paths and query parameters
- **Error Handling** - HTTP error handling patterns

Each module is a standalone Go program demonstrating specific HTTP/networking concepts.

## Requirements

- Go 1.23+

## Running Tests

```bash
cd pokedex
go test ./...
```
