package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	// Encode my user into json format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create the new put request
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// Update request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	// Create the client
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Decode the response into JSON format
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return User{}, err
	}

	return data, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// Crette a new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	// Modify the request header
	req.Header.Set("X-API-KEY", apiKey)

	// Create the client and DO
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Decode the response to JSON format
	var user User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
