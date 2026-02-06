package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(pageURL *string) (PokeApiResponse, error) {
	// Check the URL to fetch
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	if val, ok := c.cache.Get(url); ok {
		locationResponse := PokeApiResponse{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return PokeApiResponse{}, err
		}
		return locationResponse, nil
	}

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiResponse{}, err
	}

	// Get the HTTP response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiResponse{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiResponse{}, err
	}

	locationResponse := PokeApiResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return PokeApiResponse{}, err
	}

	c.cache.Add(url, data)
	return locationResponse, nil
}
