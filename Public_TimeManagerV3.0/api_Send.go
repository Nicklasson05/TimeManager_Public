package main

import (
	"fmt"
	"net/http"
)

func SendToGit(apiKey string, projectID string, duration string, issueID string, summary string) {

	url := "https://gitlab.com/api/v4/projects/" + projectID + "/issues/" + issueID + "/add_spent_time?duration=" + duration + "&summary=" + summary
	client := &http.Client{}

	// Create GET request
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
