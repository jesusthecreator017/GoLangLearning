package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
