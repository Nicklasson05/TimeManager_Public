package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var ShowTimeEntry *widget.Entry
var ShowTimeDayLabel *canvas.Text
var ShowTimeUserLabel *canvas.Text
var ShowTimeRefresh *widget.Button

func MakeShowTime() {
	ShowTimeDayLabel = canvas.NewText("// 2005-06-19 \\\\", theme.TextColor())
	ShowTimeDayLabel.TextSize = 20
	ShowTimeDayLabel.TextStyle.Bold = true
	ShowTimeDayLabel.Alignment = fyne.TextAlignCenter

	ShowTimeUserLabel = canvas.NewText("USERNAME", theme.TextColor())
	ShowTimeUserLabel.TextSize = 16
	ShowTimeUserLabel.TextStyle.Italic = true

	ShowTimeUserLabel.Alignment = fyne.TextAlignCenter

	ShowTimeEntry = widget.NewEntry()

	ShowTimeRefresh = widget.NewButtonWithIcon("Refresh", theme.MediaReplayIcon(), func() {
		LoaderActive = true
		FYCT := yct.Format("2006-01-02")
		GetTimeLogs(FYCT, GetUsername(), "")
		if TimeLogs != nil {
			ShowTimeLogs()
		} else {
			ShowTimeEntry.Text = "No TimeLogs on date"
		}
		ShowTimeDayLabel.Text = FYCT
		ShowTimeUserLabel.Text = GetUsername()
		ShowTimeDayLabel.Refresh()
		ShowTimeUserLabel.Refresh()

		ShowTimeEntry.Refresh()
		helpText("Refreshed", "blue", 1)
	})
}

type RegisterdLogs struct {
	project   string
	title     string
	iid       string
	timeSpent int
	day       string
}

var TimeLogs []RegisterdLogs

func GetTimeLogs(date string, username string, after string) {
	TimeLogs = nil

	Sdate := date + "T00:00:00Z"
	Edate := date + "T23:59:59Z"

	query := fmt.Sprintf(`

query getTimelogs{
	timelogs(startDate: "%s", endDate: "%s", username:"%s", after: "%s") {
	 nodes {
		project {
		  name
		}
		issue {
		  iid
		  title
		}
		timeSpent
		spentAt
	  }
	  count
	  pageInfo {
		hasNextPage
		hasPreviousPage
		endCursor
	  }
	}
  }
	`, Sdate, Edate, username, after)

	// Create a request body with the GraphQL query
	requestBody := map[string]string{"query": query}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request to the GitLab GraphQL API
	req, err := http.NewRequest("POST", GetGraphURL(), bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
	}

	// Set the authorization header with the personal access token
	req.Header.Set("Authorization", "Bearer "+apiKey)
	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
	}
	defer resp.Body.Close()

	// Read the response body

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
	}
	// Extracting nodes from the response
	nodes := result["data"].(map[string]interface{})["timelogs"].(map[string]interface{})["nodes"].([]interface{})

	// Iterate over each node

	for _, node := range nodes {
		project := node.(map[string]interface{})["project"].(map[string]interface{})["name"].(string)
		issue := node.(map[string]interface{})["issue"].(map[string]interface{})
		iid := issue["iid"].(string)
		title := issue["title"].(string)
		timeSpent := node.(map[string]interface{})["timeSpent"].(float64)
		spentAt := node.(map[string]interface{})["spentAt"].(string)

		TimeLogs = append(TimeLogs, RegisterdLogs{project, title, iid, int(timeSpent), spentAt})
	}
	Count := result["data"].(map[string]interface{})["timelogs"].(map[string]interface{})["count"].(float64)
	if Count != 0 {
		endCursor := result["data"].(map[string]interface{})["timelogs"].(map[string]interface{})["pageInfo"].(map[string]interface{})["endCursor"].(string)
		hasNextPage := result["data"].(map[string]interface{})["timelogs"].(map[string]interface{})["pageInfo"].(map[string]interface{})["hasNextPage"].(bool)
		if hasNextPage {
			GetTimeLogs(date, username, endCursor)
		}
	}
}

type ticket struct {
	title        string
	secondsSpent float32
}

func ShowTimeLogs() {
	var unikTask []string
	ShowTimeEntry.Text = ""
	// seperate unik tasks
	for _, v := range TimeLogs {
		unikTask = append(unikTask, v.title)
	}
	unikTask = removeDuplicates(unikTask)
	// make  structs for eatch task
	var TickerHolder []ticket
	for _, v := range unikTask {
		TickerHolder = append(TickerHolder, ticket{v, 0})
	}
	// Add time to ticket
	for index, v := range TickerHolder {
		for _, v0 := range TimeLogs {
			if v.title == v0.title {
				TickerHolder[index].secondsSpent += float32(v0.timeSpent)
			}
		}
	}
	// Create ticket look
	ShowTimeEntry.Text += "----------------Total Time Spent On Task-----------------\n"
	ShowTimeEntry.TextStyle.Bold = true
	for _, v := range TickerHolder {
		SpentTime := TranslateSeconds(int(v.secondsSpent))
		ShowTimeEntry.Text += "TimeLog: -|- Time: " + SpentTime + " -|- Task: " + v.title + "\n"
	}
	var totalTimeSpent float32
	for _, v := range TickerHolder {
		totalTimeSpent += v.secondsSpent
	}
	TotalSpentTime := TranslateSeconds(int(totalTimeSpent))
	ShowTimeEntry.Text += "TimeLog: -|- Total Time: " + TotalSpentTime + "\n"
	ShowTimeEntry.Text += "\n"
	ShowTimeEntry.Text += "\n"
	ShowTimeEntry.Text += "\n"
	ShowTimeEntry.Text += "------------------------Individual Logs---------------------\n"
	for _, v := range TimeLogs {
		SpentTime := TranslateSeconds(int(v.timeSpent))
		ShowTimeEntry.Text += "TimeLog: -|- Time: " + SpentTime + " -|- Task: " + v.title + "\n"
	}
	ShowTimeEntry.Refresh()
}

func TranslateSeconds(seconds int) string {

	var minuts float32
	var hours float32
	minuts = float32(seconds) / 60
	for minuts >= 60 {
		if minuts >= 60 {
			minuts -= 60
			hours += 1
		}
	}
	return fmt.Sprintf("%d h %d min", int(hours), int(minuts))
}
