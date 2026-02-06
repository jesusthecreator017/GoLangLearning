package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args []string) error {
	if len(args) > 1 {
		return errors.New("Too many arguments for mapf command")
	}

	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	// Update the cfg with the new pagination URLs
	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	// Display the locations
	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if len(args) > 1 {
		return errors.New("Too many arguments for mapb command")
	}

	if cfg.prevLocationURL == nil {
		return errors.New("You are on the first page!")
	}

	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	// Update the cfg with the new pagination URLs
	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	// Display the locations
	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
