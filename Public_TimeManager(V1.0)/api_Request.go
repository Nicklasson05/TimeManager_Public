package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
)

func apiRequest(selectEntry *widget.SelectEntry, apiKey string, projectID string) {
	// URL to make the GET request
	url := "https://gitlab.com/api/v4/todos"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	/*
		// Print the response body
		fmt.Println("Response:", string(body))

		// Write the response body to a JSON file
		err = ioutil.WriteFile("response.txt", body, 0644)
		if err != nil {
			fmt.Println("Error writing response to file:", err)
			return
		}
		fmt.Println("Response saved to response.txt")
	*/
	pritty := strings.Split(string(body), ",")

	for _, line := range pritty {
		//fmt.Println(line)
		writeToTasks(line, "response.txt")
	}
	readFromTasks("title", selectEntry)

	// get issue iid

	pritty0 := strings.Split(string(body), ",")

	for _, line := range pritty0 {
		//fmt.Println(line)
		writeToTasks(line, "responseIID.txt")
		//writeToTasks(line, "issueIDHolder.txt")
	}
	filterRequwstForIID("title", "\"iid\":")

	//fmt.Println(getIIDFromTask("Test Task"))

}

func writeToTasks(FormatedTime string, file string) {
	var TEMP string
	File, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()
		TEMP += line + "\n"
	}
	gg := time.Now()
	gt := gg.Format("")
	gt += FormatedTime
	TEMP += gt
	data := []byte(TEMP)
	gerr := os.WriteFile(file, data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}
}

func readFromTasks(Contains string, selectEntry *widget.SelectEntry) {
	// Clear Tasks.txt
	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile("Tasks.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	// Write Tasks to Tasks.txt
	File, err := os.Open("response.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()
		Clear := strings.Contains(line, Contains)
		if Clear {
			line = line[9:]
			line = line[:len(line)-1]
			//fmt.Printf(line + "\n")
			writeToTasks(line, "Tasks.txt")
			Clear = false
		}
	}
	// Updating selectEntry options
	var selectEntryList []string
	file, ferr := os.Open("Tasks.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		selectEntryList = append(selectEntryList, line)
		selectEntry.SetOptions(selectEntryList)
		selectEntry.Refresh()
		//fmt.Println(selectEntryList)
	}

	// Clear response.txt
	TEMP90 := ""
	data0 := []byte(TEMP90)
	qerr := os.WriteFile("response.txt", data0, 0644)
	if qerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}
}

// ID HANDERING CODE ----------------------------------------------------------------------------------
func filterRequwstForIID(title string, iid string) {
	// Clear Tasks.txt

	TEMP9 := ""
	data := []byte(TEMP9)

	gerr := os.WriteFile("issueIDHolder.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	// Write Tasks to Tasks.txt
	File, err := os.Open("responseIID.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()

		Clear0 := strings.Contains(line, iid)
		if Clear0 {
			line = line[0:]
			line = line[:]
			//fmt.Println(line + "\n")
			writeToTasks(line, "issueIDHolder.txt")
			Clear0 = false
		}

		Clear := strings.Contains(line, title)
		if Clear {
			line = line[9:]
			line = line[:len(line)-1]
			//fmt.Printf(line + "\n")
			writeToTasks(line, "issueIDHolder.txt")
			Clear = false
		}
	}
	// Clear ResponsIID

	TEMP90 := ""
	data0 := []byte(TEMP90)

	gerr0 := os.WriteFile("responseIID.txt", data0, 0644)
	if gerr0 != nil {
		fmt.Println("Error while writeing to Tasks")
	}
}

func getIIDFromTask(task string) string {
	var lineHolder string
	File, err := os.Open("issueIDHolder.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()

		Clear01 := strings.Contains(line, task)
		if Clear01 {
			//line = line[0:]
			//line = line[:]
			//fmt.Println(line + "\n")
			return lineHolder
		}

		Clear0 := strings.Contains(line, "\"iid\":")
		if Clear0 {
			line = line[6:]
			//line = line[:]
			lineHolder = line
			//fmt.Println(line + "\n")
			Clear0 = false
		}
	}
	return "Error"
}
