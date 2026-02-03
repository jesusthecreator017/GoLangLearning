package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	// ?
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	return resp.StatusCode, nil
}
