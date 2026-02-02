package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// Encode the user as a json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// Set the header of the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	// Create a new client and make the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON data from the response into a new User struct
	var user User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
