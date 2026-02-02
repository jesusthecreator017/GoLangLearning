package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// Create the new request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Update request header
	req.Header.Set("X-API-KEY", apiKey)

	// Create client and process the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 299 {
		return fmt.Errorf("request to user was unsuccessful")
	}
	fmt.Println("requesy to delete user was Successful")

	return nil
}
