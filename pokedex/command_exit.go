package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(state *config, args []string) error {
	if len(args) > 1 {
		return errors.New("Too many arguments for exit command")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
