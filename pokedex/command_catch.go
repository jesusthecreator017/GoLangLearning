package main

import (
	"errors"
)

func commandCatch(cfg *config, args []string) error {
	// Check if the pokemon name has been passed into the command
	if len(args) < 2 {
		return errors.New("Please provide a pokemon name to catch.")
	}

	// For now, lets print out the pokemon name.
	return nil
}
