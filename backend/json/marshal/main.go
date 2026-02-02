package main

import (
	"encoding/json"
	"fmt"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var res [][]byte
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("Error when converting to marshal: %w", err)
		}
		res = append(res, data)
	}

	return res, nil
}
