package main

import (
	"bufio"
	"os"
	"strings"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Send To Git
var GitProgressBar *widget.ProgressBar
var GitCheckMark *widget.Check
var PjNavNeg *widget.Check
var PjLable *widget.Select

var readyToSendLog []string
var sendToGitEntry *widget.Entry
var sendToGitButton *widget.Button

func MakeSendToGit() {
	readyToSendLog = []string{}
	restrictMarkIcon = widget.NewIcon(theme.CancelIcon())
	sendToGitEntry = widget.NewMultiLineEntry()
	sendToGitEntry.PlaceHolder = "Mark dates to get started"

	GitProgressBar = widget.NewProgressBar()

	GitCheckMark = widget.NewCheck("Show Log", func(b bool) {
		if b {
			entry.Show()
			sendToGitEntry.Hide()
		} else {
			entry.Hide()
			sendToGitEntry.Show()
			loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
			GetNumOfTasks()
			sendToGitEntry.Refresh()
		}
	})

	// UPLOAD BUTTON
	sendToGitButton = widget.NewButtonWithIcon("Send", theme.UploadIcon(), func() {
		// Send to git
		if CheckAPIKEY(apiKey) {

			ReadyTasksForGit()
			// clean up after posting to gitlab
			if uploadComplete {
				//save datets that cannot be upploaded again
				for _, v := range readyToSendLog {
					writeToTasks(v, Directory+"/markedDates.txt")
				}

				// clear dates
				readyToSendLog = nil

				FormatedYCT10 := yct.Format("(2006-01-02)")
				clearDate := fileContents(Directory+"/markedDates.txt", FormatedYCT10)
				if clearDate {
					restrictMarkIcon.Show()
					markButton.Disable()
				} else {
					restrictMarkIcon.Hide()
					markButton.Enable()
				}
				sendToGitEntry.Text = "Press send button to log time to gitlab"
				sendToGitEntry.Refresh()
				helpText("Upload", "blue", 2)
				GitProgressBar.Value = 0
				GitProgressBar.Refresh()
			}
		} else {
			helpText("Enter valid apikey", "yellow", 2)
		}
	})
}

func GetTaskIDS(TaskName string) (string, string) {
	logFile, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()

		if strings.Contains(line, "-|- Task:") {
			if strings.Contains(line, TaskName) {
				if strings.Contains(line, "-|-iid:") {
					id := takeTextFromString(line, strings.Index(line, "-|-iid:")+7, strings.Index(line, "-|-PJID:"))
					pjid := takeTextFromString(line, strings.Index(line, "-|-PJID:")+8, len(line))
					return id, pjid
				}
			}
		}
	}
	//return GetID(TaskName)
	return GetID(TaskName), GetPJID(TaskName)
}
