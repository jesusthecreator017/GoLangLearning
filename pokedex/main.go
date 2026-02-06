package main

import (
	"time"

	"learning.com/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.PokeApiPokemonInfoResponse),
	}
	startRepl(cfg)
}
