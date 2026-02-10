package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("Your pokedex is empty! Try catching some pokemon first!")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
