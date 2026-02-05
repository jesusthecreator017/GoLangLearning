package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokeApiResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []mapData `json:"results"`
}

type mapData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(state *config) error {
	// Get the map data from the PokeAPI
	url := "https://pokeapi.co/api/v2/location-area/"
	data, err := getPokeApiMapData(url)
	if err != nil {
		return fmt.Errorf("Error fetching map data: %w", err)
	}

	// Update the state with the next URL
	state.Next = data.Next
	state.Prev = data.Previous

	// Display the map data
	for _, value := range data.Results {
		fmt.Printf("%v\n", value.Name)
	}

	return nil
}

func getPokeApiMapData(url string) (*pokeApiResponse, error) {
	// Fetch the map data from the PokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %w", err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var data pokeApiResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Error decoding response: %w", err)
	}

	return &data, nil
}
