package main

func fetchTasks(baseURL, availability string) []Issue {
	fullURL := baseURL + "?sort=estimate"
	if availability != "" {
		switch availability {
		case "Low":
			fullURL += "&limit=1"
		case "Medium":
			fullURL += "&limit=3"
		case "High":
			fullURL += "&limit=5"
		default:
			break
		}
	}
	return getIssues(fullURL)
}
