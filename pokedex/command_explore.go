package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	// Check if the area has been passed into the command
	if len(args) < 2 {
		return errors.New("Please provide a location name to explore.")
	}

	// Fetch the location info
	locationResp, err := cfg.pokeapiClient.GetLocationInfoList(&args[1])
	if err != nil {
		return errors.New("Failed to retrieve location info.")
	}

	// Display just the pokemon encounters for now
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}

	return nil
}
