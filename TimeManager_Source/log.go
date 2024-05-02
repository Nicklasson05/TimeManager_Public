package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Log
var selectEntry *widget.Select
var summaryEntry *widget.Entry
var PlayButton *widget.Button
var RecordButton *widget.Button
var image *widget.Button
var refeshButton *widget.Button
var entry *widget.Entry

// LOG MENU TIME HOLDERS
var arrActive bool = false
var arrivedTime time.Time
var latestArchivedTime time.Time
var firstTimeArchivedTime bool = true

var sidebar *widget.Toolbar
var selectEntryList []string
var temp []string

func MakeLog() {

	temp = []string{
		"abbe",
	}
	ProjectIssuesHolder = append(ProjectIssuesHolder, ProjectMap{1, "Temp", temp})
	entry = widget.NewMultiLineEntry()
	selectEntryList = []string{"Enter a valid API key in options"}
	selectEntry = widget.NewSelect(selectEntryList, func(s string) {

	})
	summaryEntry = widget.NewEntry()

	PlayButton = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		if !arrActive {
			//Navigate to current day

			// LOG ARRIVED
			arrActive = true
			currentTime := time.Now()
			NavigateTooCurrentDay(currentTime, date, entry, readyToSendLog, markButton, restrictMarkIcon)
			arrivedTime = currentTime
			latestArchivedTime = arrivedTime
			entry.TextStyle = fyne.TextStyle{Bold: true}
			FormatedTime := currentTime.Format("( 15:04 ) ")
			FormatedTime += "-|- Arrived.\n "
			entry.Text += FormatedTime
			entry.MultiLine = true
			entry.Refresh()

			// GORUTINE

			if !liveClockActive && !liveClockOverrideActive {
				ChangeEndTime()
				//go liveClock(c, timeDoneText, quit, endTime, arrivedTime, timeLeftText, latestArchivedTime)
				go LiveClockOverride(ch, quit0, "Arrived", shiftLeft, endTime, timeLeftText, arrivedTime, timeDoneText)
				liveClockActive = true
			}

			// WRITE TO LOG
			writeToLog(FormatedTime)

			//Reset status texxts based on the new day
			ResetStatusTexts(ch, quit0, quit, true)
		}

	})

	RecordButton = widget.NewButtonWithIcon("", theme.MediaRecordIcon(), func() {
		if arrActive {
			if selectEntry.Selected != "" {
				if restrictMarkIcon.Hidden {
					if ShowWebUrlBool {
						for _, v := range ProjectIssuesHolder {
							if v.ProjectName == PjLable.Selected {
								for _, v0 := range v.Tasks {
									if strings.Contains(selectEntry.Selected, v0) {
										selectEntry.Selected = v0
									}
								}
							}
						}
					}
					currentTime := time.Now()
					if firstTimeArchivedTime {
						latestArchivedTimeZoneFix := latestArchivedTime.Add(time.Duration(timeModifier) * time.Hour)
						timeWork += currentTime.Sub(latestArchivedTimeZoneFix)
						firstTimeArchivedTime = false
					} else {
						timeWork += currentTime.Sub(latestArchivedTime)
					}

					TEMP := fmt.Sprintf("%d h %d min", int(timeWork.Hours()), int(timeWork.Minutes())%60)
					timeWorkText.Text = "Work Done: " + TEMP
					timeWorkText.Refresh()

					entry.TextStyle = fyne.TextStyle{Bold: true}
					FormatedTime := latestArchivedTime.Format("( 15:04 - ")
					FormatedTime += currentTime.Format("15:04 )")

					FormatedTime += " -|- Task: "
					FormatedTime += selectEntry.Selected
					FormatedTime += " -|- Summary: "
					FormatedTime += summaryEntry.Text
					LogVer := FormatedTime
					FormatedTime += " \n"
					entry.Text += FormatedTime

					latestArchivedTime = currentTime
					entry.MultiLine = true
					TaskName := selectEntry.Selected
					selectEntry.Selected = ""
					selectEntry.FocusLost()

					summaryEntry.Text = ""
					summaryEntry.FocusLost()

					entry.Refresh()
					LogVer += "-|-iid:" + GetID(TaskName) + "-|-PJID:" + GetPJID(TaskName) + " \n"

					writeToLog(LogVer)

					FcurrentTime := time.Now().Format("(2006-01-02)")
					search := yct.Format("(2006-01-02)")
					turnOffClock = true
					if FcurrentTime != search {
						ResetStatusTexts(ch, quit0, quit, false)
					} else {
						ResetStatusTexts(ch, quit0, quit, true)
					}

				} else {
					helpText("Restricted date", "red", 1)
				}

			} else {
				helpText("Select a Task", "yellow", 1)
			}
		} else {
			if !arrActive {
				helpText("Arrived Needed", "yellow", 1)
			}
		}
	})

	sidebar = widget.NewToolbar()
	//Project Nav
	PjLable = widget.NewSelect(ProjectList, func(s string) {

	})
	PjLable.PlaceHolder = "Select Project"

	PjLable.OnChanged = func(input string) {
		for index, issue := range ProjectIssuesHolder {
			if input == issue.ProjectName {
				if ShowWebUrlBool {
					ShowWebUrl()
				} else {
					selectEntry.SetOptions(ProjectIssuesHolder[index].Tasks)
				}
			}
		}
	}

	refeshButton = widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
		LoaderActive = true
		if CheckAPIKEY(apiKey) {
			userID := RequestUserID(apiKey)
			RequestIssues(apiKey, userID)
			FoundPJ := false
			for index, issue := range ProjectIssuesHolder {
				if PjLable.Selected == issue.ProjectName {
					selectEntry.SetOptions(ProjectIssuesHolder[index].Tasks)
					FoundPJ = true
				}
			}
			if !FoundPJ {
				PjLable.Selected = ""
				var emptyList []string
				selectEntry.SetOptions(emptyList)
				PjLable.Refresh()
			}

			helpText("Tasks Refreshed", "blue", 1)
		} else {
			helpText("Enter valid apikey", "yellow", 2)
		}
	})

	selectEntry.PlaceHolder = "Select Task"
	selectEntry.Refresh()

	summaryEntry.PlaceHolder = "Summary"
	summaryEntry.Refresh()

	// MAIN ENTRY SAVE BUTTON
	image = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		LoaderActive = true
		if entry.Text == "" {
			helpText("ERROR", "red", 1)
			restart()
		}
		// Remove last disrutbing charectors " " && "\n"
		for {
			lastChar := entry.Text[len(entry.Text)-1:]
			if lastChar == " " || lastChar == "\n" {
				entry.Text = entry.Text[:len(entry.Text)-1]
			} else {
				break
			}
		}
		// Split entry.Text into more strings
		hh := strings.Split(entry.Text, "\n")
		// Remove ':' and spaces
		var position int = -1
		charector := '('
		var hf []string

		for _, value := range hh {
			for i, char := range value {
				if char == charector {
					position = i
					break
				}
			}
			if position != -1 {
				value = value[position:]
				value = value[:]
				value = LogSaveIDS(value)
				hf = append(hf, value)
			}
		}

		FetchLegacyTasks()
		if ValidateTask(hf) {

			var hn []string
			FYCT := yct.Format("(2006-01-02)")
			for _, v := range hf {
				hn = append(hn, FYCT+" "+v)
			}
			// Remove date from log
			removeLineFromFile(Directory+"/log.txt", FYCT)
			// Write new date infro into log
			for _, v := range hn {
				newWriteToLog(v)
			}
			entry.Text += "\n"
			readFromLog(yct, entry)
			//restart()

			FcurrentTime := time.Now().Format("(2006-01-02)")
			turnOffClock = true
			if FcurrentTime != FYCT {
				ResetStatusTexts(ch, quit0, quit, false)
			} else {
				ResetStatusTexts(ch, quit0, quit, true)
			}

			// SAVE STATUS TEXT
			helpText("SAVED", "green", 1)
		} else {
			helpText("Invalid Task name", "red", 2)
		}
	})
}

func LoadInLog() {

	tnf = yct.Format("(2006-01-02)")

	logFile, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		Clear := strings.Contains(line, tnf)
		if Clear {
			new := strings.Trim(line, tnf)
			// fix arrived time
			arrivedTime = updateArrivedTime(line, arrivedTime)
			ChangeEndTime()
			// fix latest time
			latestArchivedTime = updateLatestArchivedTime(line, latestArchivedTime)

			// status text
			shiftLeft = updateShiftLeftText(endTime)
			workTime += updateWorkDoneStatus(line, totalWorkTime)
			arr := strings.Contains(line, "Arrived")
			if arr {
				arrActive = true
			}

			if !liveClockActive && !liveClockOverrideActive {
				go LiveClockOverride(ch, quit0, line, shiftLeft, endTime, timeLeftText, arrivedTime, timeDoneText)
				liveClockOverrideActive = true
			}
			// Write to Entry
			entry.MultiLine = true
			entry.Text += new + "\n"
			entry.Refresh()
		}

	}
	timeWork = workTime

	TEMP := "Work Done: "
	TEMP += fmt.Sprintf("%d h %d min", int(workTime.Hours()), int(workTime.Minutes())%60)
	timeWorkText.Text = TEMP
	timeWorkText.Refresh()

	TEMP13 := "Shift Left: "
	TEMP13 += fmt.Sprintf("%d h %d min ", int(shiftLeft.Hours()), int(shiftLeft.Minutes())%60)
	timeLeftText.Text = TEMP13
	timeLeftText.Refresh()

	readFromLog(yct, entry)
}

func LogSaveIDS(value string) string {
	if strings.Contains(value, "-|- Task:") {

		VTask := takeTextFromString(value, strings.Index(value, "-|- Task:")+9, strings.Index(value, "-|- Summary:"))

		logFile, err := os.Open(Directory + "/log.txt")
		if err != nil {
			panic(err)
		}

		logScanner := bufio.NewScanner(logFile)
		for logScanner.Scan() {
			line := logScanner.Text()

			if strings.Contains(line, "-|- Task:") {
				LTask := takeTextFromString(line, strings.Index(line, "-|- Task:")+9, strings.Index(line, "-|- Summary:"))
				if LTask == VTask {
					if strings.Contains(line, "-|-iid:") {
						idstring := takeTextFromString(line, strings.Index(line, "-|-iid:"), len(line))
						value += idstring
						return value
					}
				}
			}
		}
	}
	return value
}
