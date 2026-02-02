package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://1.1.1.1/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response body: %w", err)
	}

	var dnsResponse DNSResponse
	if err := json.Unmarshal(data, &dnsResponse); err != nil {
		return "", fmt.Errorf("error when unmarhsalling response body: %w", err)
	}

	if len(dnsResponse.Answer) == 0 {
		return "", fmt.Errorf("no A records found for %s", domain)
	}

	return dnsResponse.Answer[0].Data, nil
}
