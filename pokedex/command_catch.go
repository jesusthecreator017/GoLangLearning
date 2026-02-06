package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	// Check if the pokemon name has been passed into the command
	if len(args) < 2 {
		return errors.New("Please provide a pokemon name to catch.")
	}

	// Fetch the pokemon info
	pokemonResp, err := cfg.pokeapiClient.GetPokemonInfoList(&args[1])
	if err != nil {
		return errors.New("Failed to retrieve pokemon info.")
	}

	// Handle catching logic: Use the pokemons base experience as a proxy for catch chance
	catchAttemp := rand.Float64()
	catchChance := 1.0 - float64(pokemonResp.BaseExperience)/300.0 // Normalize to a value between 0 and 1

	if catchChance < 0.05 {
		catchChance = 0.05 // Minimum catch chance of 5%
	}

	fmt.Printf("Throwing a pokeball at %s...\n", pokemonResp.Name)
	time.Sleep(2 * time.Second)

	if catchAttemp <= catchChance {
		// Successful catch
		cfg.pokedex[pokemonResp.Name] = pokemonResp
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}

	return nil
}
