package main

import (
	"fmt"
	"net/http"
)

func SendToGit(apiKey string, projectID string, duration string, issueID string) {
	// URL to make the GET request
	url := "https://gitlab.com/api/v4/projects/" + projectID + "/issues/" + issueID + "/add_spent_time?duration=" + duration + ""
	//url := "https://gitlab.com/api/v4/projects/54755108/issues/7/add_spent_time?duration=1h"
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("PRIVATE-TOKEN", apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
}
