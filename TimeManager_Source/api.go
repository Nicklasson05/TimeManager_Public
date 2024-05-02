package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Issue struct {
	ID     int    `json:"iid"`
	Title  string `json:"title"`
	PJID   int    `json:"project_id"`
	WebUrl string `json:"web_url"`
}

type PJFilter struct {
	PJName string `json:"name"`
}

type ResFilter struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

type stringStruct struct {
	Duration, Task, Summary string
}

type ProjectMap struct {
	ProjectID   int
	ProjectName string
	Tasks       []string
}

var ProjectIssuesHolder []ProjectMap
var ProjectList []string

var URL string
var apiKey string

var issueHolder []Issue
var PostInfo []stringStruct

var uploadComplete bool

// post times to git issues
func SendToGit0(apiKey string, projectID string, duration string, issueID string, summary string) {
	if projectID == "Error" {
		helpText("Invalid APIKey", "red", 3)
		return
	}
	if issueID == "Error" {
		helpText("Invalid APIKey", "red", 3)
		return
	}
	url := URL + "/projects/" + projectID + "/issues/" + issueID + "/add_spent_time?duration=" + duration + "&summary=" + summary
	client := &http.Client{}
	// Create GET request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		helpText("Upload Failed", "red", 3)
		return
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("PRIVATE-TOKEN", apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		helpText("Upload Failed", "red", 3)
		return
	}
	defer resp.Body.Close()
	uploadComplete = true
}

// get issues
func RequestIssues(apiKey string, userID string) {
	// Clear respons for request
	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile(Directory+"/response.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	// setup url
	url := URL
	url += "/issues?assignee_id=" + userID
	url += "&issue_type=issue&scope=assigned_to_me&state=opened&per_page=100"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Error Timeout check so that u have a connection to your server or that your url is correctly inputted")
		} else {
			fmt.Println("Error sending request:", err)
		}
	}
	defer resp.Body.Close()

	// Filter response for issue ID, Title, Project ID
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		fmt.Println(err)
	}
	ProjectIssuesHolder = nil
	issueHolder = nil
	for _, issue := range issues {
		issueHolder = append(issueHolder, Issue{ID: issue.ID, Title: issue.Title, PJID: issue.PJID, WebUrl: issue.WebUrl})
	}
	ProjectList = nil
	FilterIssuesByProject()
	UpdateSelectProjectOptions()
}

// get userID
func RequestUserID(apiKey string) string {

	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile(Directory+"/response.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	url := URL
	url += "/user"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// GET TASKS TO Tasks.txt
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	// Write to Response
	pritty := strings.Split(string(body), ",")
	for _, line := range pritty {
		if strings.Contains(line, "\"id\":") {
			line = line[6:]
			UserID = line
			return UserID
		}
	}
	return "Error"
}

// get issue project id
func GetPJID(task string) string {
	for _, element := range issueHolder {

		if strings.Contains(element.Title, task) {
			FPJID := strconv.Itoa(element.PJID)
			return FPJID
		}
	}
	return "Error"
}

// get username
func GetUsername() string {
	url := URL
	url += "/user"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	// GET TASKS TO Tasks.txt
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	// Write to Response
	pritty := strings.Split(string(body), ",")
	for _, line := range pritty {
		if strings.Contains(line, "\"username\":") {
			return takeTextFromString(line, 12, len(line)-1)
		}
	}
	return "Error"
}

// get issue iid
func GetID(task string) string {
	for _, element := range issueHolder {
		if element.Title == task {
			FID := strconv.Itoa(element.ID)
			return FID
		}
	}
	return "Error"
}

// check is apikey is valid
func CheckAPIKEY(apiKey string) bool {
	url := URL
	url += "/version"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	if strings.Contains(string(body), "\"version\"") {
		return true
	} else {
		return false
	}
}

// Get project name from PJID
func GetPJname(PJID string) string {

	// setup url
	url := URL
	url += "/projects/" + PJID

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Error Timeout check so that u have a connection to your server or that your url is correctly inputted")
		} else {
			fmt.Println("Error sending request:", err)
		}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Filter response for issue ID, Title, Project ID
	var PJfilter PJFilter
	err = json.Unmarshal(body, &PJfilter)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return PJfilter.PJName
}

// Get Project namespace path

// OTHER
func ReadyTasksForGit() {
	var DurationHolder string
	var TaskHolder string
	var SummaryHolder string
	var currentDate string

	File, err := os.Open(Directory + "/response.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()
		clear2 := strings.Contains(line, "-|- Date: ")
		if clear2 {
			currentDate = takeTextFromString(line, 11, len(line)-1)
		}
		clear0 := strings.Contains(line, "-|- Summary:")
		if clear0 {
			// Find Summary
			NumofTasks += 1
			index := strings.LastIndex(line, "-|- Summary:")
			summary := takeTextFromString(line, index+12, len(line))
			SummaryHolder = summary
			// Post
			PostInfo = append(PostInfo, stringStruct{DurationHolder, TaskHolder, SummaryHolder})
			iid, pjid := GetTaskIDS(TaskHolder)
			SendToGit(pjid, DurationHolder, iid, replaceSpaces(SummaryHolder), currentDate)
			GitProgressBar.Value += TaskTakeValue
			GitProgressBar.Refresh()
		}
		clear := strings.Contains(line, "-->")
		if clear {
			// Find Duration
			duration := takeTextFromString(line, 0, 5)
			DurationHolder = duration
			// Find Task
			index := strings.LastIndex(line, "-->")
			task := takeTextFromString(line, index+11, len(line))
			TaskHolder = task
		}
	}
}

// Experimental
func SendToGit(projectID string, duration string, issueID string, summary string, spentAt string) {
	// Replace with your GitLab personal access token
	issueGID := GetIssueGID(issueID, projectID)
	// Define the GraphQL query to fetch issues
	query := fmt.Sprintf(`
	mutation {
		timelogCreate(input: {
			issuableId: "%s",
			timeSpent: "%s",
			spentAt: "%s"
			summary: "%s"
		}){
			timelog {
				id
				timeSpent
			}
		}
	}
	`, issueGID, duration, spentAt, summary)

	// Create a request body with the GraphQL query
	requestBody := map[string]string{"query": query}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		helpText("Upload Failed", "red", 3)
		fmt.Println("Error marshalling request body:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request to the GitLab GraphQL API
	req, err := http.NewRequest("POST", GetGraphURL(), bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		helpText("Upload Failed", "red", 3)
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the authorization header with the personal access token
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		helpText("Upload Failed", "red", 3)
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	uploadComplete = true
}

func GetIssueGID(issueIID string, ProjectID string) string {

	FissueIID := "\"" + issueIID + "\""
	FProjectID := GetPJNamespace(ProjectID)

	query := fmt.Sprintf(`
		query {
			project(fullPath: "%s") {
				issue(iid: %s) {
					id
				}
			}
		}
	`, FProjectID, FissueIID)

	// Create a request body with the GraphQL query
	requestBody := map[string]string{"query": query}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return "Error1"
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request to the GitLab GraphQL API
	req, err := http.NewRequest("POST", GetGraphURL(), bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return "Error2"
	}

	// Set the authorization header with the personal access token
	req.Header.Set("Authorization", "Bearer "+apiKey)
	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return "Error3"
	}
	defer resp.Body.Close()

	// Read the response body

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return "Error4"
	}

	issue := result["data"].(map[string]interface{})["project"].(map[string]interface{})["issue"].(interface{})

	id := issue.(map[string]interface{})["id"].(string)
	return id

}

// Get Project namespace path
type PJNSPfilter struct {
	PJNameSpace string `json:"path_with_namespace"`
	PJID        int    `json:"id"`
}

func GetPJNamespace(projectID string) string {

	// setup url
	url := URL
	url += "/projects/?membership=true"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Error Timeout check so that u have a connection to your server or that your url is correctly inputted")
		} else {
			fmt.Println("Error sending request:", err)
		}
	}
	defer resp.Body.Close()

	var PJNSP []PJNSPfilter
	if err := json.NewDecoder(resp.Body).Decode(&PJNSP); err != nil {
		fmt.Println(err)
	}

	FPJID, _ := strconv.Atoi(projectID)

	for _, v := range PJNSP {
		if v.PJID == FPJID {
			return v.PJNameSpace
		}
	}

	return "Error"
}

func GetGraphURL() string {
	return takeTextFromString(URL, 0, strings.Index(URL, "/api/")+5) + "graphql"
}
