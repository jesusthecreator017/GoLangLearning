package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	// Check if the pokemon name has been passed into the command
	if len(args) < 2 {
		return errors.New("Please provide a pokemon name to inspect.")
	}

	// Check if the pokemon exists in the pokedex
	pokemonName := args[1]
	pokemonInfo, exists := cfg.pokedex[pokemonName]
	if !exists {
		return fmt.Errorf("Pokemon '%s' not found in your pokedex. Try catching it first!", pokemonName)
	}

	// Display the pokemon information
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemonInfo.Types {
		fmt.Printf(" - %s\n", typeInfo.Type.Name)
	}

	return nil
}
