package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	if len(args) > 1 {
		return errors.New("Too many arguments for help command")
	}

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
