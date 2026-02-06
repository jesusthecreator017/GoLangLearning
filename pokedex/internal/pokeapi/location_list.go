package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(pageURL *string) (PokeApiLocationResponse, error) {
	// Check the URL to fetch
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	if val, ok := c.cache.Get(url); ok {
		locationResponse := PokeApiLocationResponse{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return PokeApiLocationResponse{}, err
		}
		return locationResponse, nil
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	// Get the HTTP response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	locationResponse := PokeApiLocationResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return PokeApiLocationResponse{}, err
	}

	c.cache.Add(url, data)
	return locationResponse, nil
}

func (c *Client) GetLocationInfoList(areaUrl *string) (PokeApiLocationInfoResponse, error) {
	// Check the URL to fetch
	url := baseURL + "/location-area/" + *areaUrl

	// Check the cache first
	if val, ok := c.cache.Get(url); ok {
		locationInfoResponse := PokeApiLocationInfoResponse{}
		err := json.Unmarshal(val, &locationInfoResponse)
		if err != nil {
			return PokeApiLocationInfoResponse{}, err
		}
		return locationInfoResponse, nil
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiLocationInfoResponse{}, err
	}

	// Get the HTTP response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiLocationInfoResponse{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiLocationInfoResponse{}, err
	}

	locationInfoResponse := PokeApiLocationInfoResponse{}
	err = json.Unmarshal(data, &locationInfoResponse)
	if err != nil {
		return PokeApiLocationInfoResponse{}, err
	}

	c.cache.Add(url, data)
	return locationInfoResponse, nil
}

func (c *Client) GetPokemonInfoList(pokemonUrl *string) (PokeApiPokemonInfoResponse, error) {
	// Check the URL to fetch
	url := baseURL + "/pokemon/" + *pokemonUrl

	// Check the cache first
	if val, ok := c.cache.Get(url); ok {
		pokemonInfoResponse := PokeApiPokemonInfoResponse{}
		err := json.Unmarshal(val, &pokemonInfoResponse)
		if err != nil {
			return PokeApiPokemonInfoResponse{}, err
		}
		return pokemonInfoResponse, nil
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiPokemonInfoResponse{}, err
	}

	// Get the HTTP response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiPokemonInfoResponse{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiPokemonInfoResponse{}, err
	}

	pokemonInfoResponse := PokeApiPokemonInfoResponse{}
	err = json.Unmarshal(data, &pokemonInfoResponse)
	if err != nil {
		return PokeApiPokemonInfoResponse{}, err
	}

	c.cache.Add(url, data)
	return pokemonInfoResponse, nil
}
