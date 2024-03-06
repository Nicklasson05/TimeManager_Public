package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeGUI() fyne.CanvasObject {

	// FETCH APPLICATIONS INFO FROM FILES
	var endTime time.Time
	var err error
	var apiKey string
	var projectID string

	file, ferr := os.Open("useroptions.txt")
	if ferr != nil {
		panic(ferr)
	}
	var i int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		// Load endtime option
		if i == 1 {
			endTime, err = time.Parse("15:04", line)
			if err != nil {
				fmt.Println("Error while reading from useroptions.txt")
			}
		}
		// Load API KEY
		if i == 2 {
			apiKey = line
		}

		if i == 3 {
			projectID = line
		}
	}

	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile("response.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	// DECLARITIONER OF VARUABELS

	// UI ELEMENTS
	content := canvas.NewRectangle(color.Gray{Y: 0xee})
	entry := widget.NewMultiLineEntry()
	selectEntryList := []string{"Enter a valid API key in options"}
	selectEntry := widget.NewSelectEntry(selectEntryList)
	summaryEntry := widget.NewEntry()
	readyToSendLog := []string{}
	restrictMarkIcon := widget.NewIcon(theme.CancelIcon())
	sendToGitEntry := widget.NewMultiLineEntry()
	sendToGitEntry.PlaceHolder = "Mark dates that will be sent"

	// API RELATED VARUABELS
	apiRequest(selectEntry, apiKey, projectID)
	apiKeyLabel := widget.NewLabel("Enter your gitlab API key here")
	apiKeyEntry := widget.NewPasswordEntry()

	// PROJECTID OPTIONS
	projectIDLabel := canvas.NewText("", color.RGBA{0, 255, 0, 155}) // NOW IS SAVED STATUS TEXT
	projectIDLabel.TextStyle.Bold = true
	projectIDEntry := widget.NewLabel("")

	// DATE TEXT VARUABELS
	date := widget.NewLabel("")
	ct := time.Now()
	date.Text = ct.Weekday().String()
	date.Text += ct.Format(", 2006-01-02 ")

	// CLOCK GOROUTINE
	yct := ct
	c := make(chan int)
	quit := make(chan bool)

	ch := make(chan int)
	quit0 := make(chan bool)

	// CLOCK STATUS HOLDERS
	var liveClockActive bool
	var liveClockOverrideActive bool

	// LOG MENU TIME HOLDERS
	var arrivedTime time.Time
	var latestArchivedTime time.Time
	var firstTimeArchivedTime bool = true

	// One per day restriction section
	var eodActive bool = false
	var arrActive bool = false

	// STATUS TEXTS VALUE HOLDERS
	var timeSlacking time.Duration
	var timeWork time.Duration
	var timeDone time.Duration
	var shiftLeft time.Duration

	// STATUS TEXTS

	// slacking
	timeStatText := canvas.NewText("Slacking: 0 h 0 min ", color.Black)
	timeStatText.TextSize = 15
	// WorkDone
	timeWorkText := canvas.NewText("Work Done: 0 h 0 min ", color.Black)
	timeWorkText.TextSize = 15
	// shift left
	timeLeftText := canvas.NewText("Shift Left: 0 h 0 min ", color.Black)
	timeLeftText.TextSize = 15
	// worked today
	timeDoneText := canvas.NewText("Worked Today: 0 h 0 min", color.Black)
	timeDoneText.TextSize = 15

	// MARK BUTTON CODE

	markButton := widget.NewCheck("", func(b bool) {
		FormatedYCT10 := yct.Format("(2006-01-02)")
		clearDate := fileContents("markedDates.txt", FormatedYCT10)
		if !clearDate {
			if b {
				FormatedYCT := yct.Format("(2006-01-02)")
				readyToSendLog = append(readyToSendLog, FormatedYCT)
			} else if !b {
				FormatedYCT0 := yct.Format("(2006-01-02)")
				readyToSendLog = removeElementFromList(readyToSendLog, FormatedYCT0)
				readyToSendLog = removeDuplicates(readyToSendLog)
			}
		}
	})

	// DATE NAVIGATION
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			yct = yct.Add(-24 * time.Hour)
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()
			readFromLog(yct, entry)

			// Check Box
			var exists bool
			search := yct.Format("(2006-01-02)")
			for _, v := range readyToSendLog {
				if v == search {
					exists = true
					break
				}
			}
			if exists {
				markButton.SetChecked(true)
				readyToSendLog = removeDuplicates(readyToSendLog)
			} else if !exists {
				markButton.SetChecked(false)
			}

			FormatedYCT10 := yct.Format("(2006-01-02)")
			clearDate := fileContents("markedDates.txt", FormatedYCT10)
			if clearDate {
				restrictMarkIcon.Show()
				markButton.Disable()
			} else {
				restrictMarkIcon.Hide()
				markButton.Enable()
			}
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			yct = yct.Add(+24 * time.Hour)
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()

			readFromLog(yct, entry)

			// Check Box
			var exists bool
			search := yct.Format("(2006-01-02)")
			for _, v := range readyToSendLog {
				if v == search {
					exists = true
					break
				}
			}
			if exists {
				markButton.SetChecked(true)
				readyToSendLog = removeDuplicates(readyToSendLog)
			} else if !exists {
				markButton.SetChecked(false)
			}

			FormatedYCT10 := yct.Format("(2006-01-02)")
			clearDate := fileContents("markedDates.txt", FormatedYCT10)
			if clearDate {
				restrictMarkIcon.Show()
				markButton.Disable()
			} else {
				restrictMarkIcon.Hide()
				markButton.Enable()
			}

		}),
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			yct = time.Now()
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()

			readFromLog(yct, entry)

			// Check Box
			var exists bool
			search := yct.Format("(2006-01-02)")
			for _, v := range readyToSendLog {
				if v == search {
					exists = true
					break
				}
			}
			if exists {
				markButton.SetChecked(true)
				readyToSendLog = removeDuplicates(readyToSendLog)
			} else if !exists {
				markButton.SetChecked(false)
			}

			FormatedYCT10 := yct.Format("(2006-01-02)")
			clearDate := fileContents("markedDates.txt", FormatedYCT10)
			if clearDate {
				restrictMarkIcon.Show()
				markButton.Disable()
			} else {
				restrictMarkIcon.Hide()
				markButton.Enable()
			}

		}),
	)

	// UPLOAD BUTTON
	sendToGitButton := widget.NewButtonWithIcon("", theme.UploadIcon(), func() {
		fmt.Println("UpLoad Proceding")
		File, err := os.Open("response.txt")
		if err != nil {
			panic(err)
		}
		Scanner := bufio.NewScanner(File)
		for Scanner.Scan() {
			line := Scanner.Text()
			clear := strings.Contains(line, "-->")
			if clear {
				duration := takeTextFromString(line, 0, 5)
				fmt.Println(duration)
				task := line
				task = task[16:]
				//Remove Last Spaces
				task = task[:len(task)-2]
				//Remove first Space
				firstChar := task[0]
				firstCharString := string(firstChar)
				if firstCharString == " " {
					task = task[1:]
				}
				// API POST
				SendToGit(apiKey, GetPJID(task), duration, getIIDFromTask(task), GetSummaryFromTask(task))
			}

		}

		// clean up after posting to gitlab

		//save datets that cannot be upploaded again
		for _, v := range readyToSendLog {
			writeToTasks(v, "markedDates.txt")
		}

		// clear dates
		readyToSendLog = nil
		fmt.Println("reoSendLog: Nil")
		// show completion
		sendToGitEntry.Text = "Upload Complete..."
		sendToGitEntry.Refresh()

		FormatedYCT10 := yct.Format("(2006-01-02)")
		clearDate := fileContents("markedDates.txt", FormatedYCT10)
		if clearDate {
			restrictMarkIcon.Show()
			markButton.Disable()
		} else {
			restrictMarkIcon.Hide()
			markButton.Enable()
		}

	})

	// LOG MENU
	sidebar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			if !arrActive {
				// LOG ARRIVED
				arrActive = true
				currentTime := time.Now()
				arrivedTime = currentTime
				latestArchivedTime = arrivedTime
				entry.TextStyle = fyne.TextStyle{Bold: true}
				FormatedTime := currentTime.Format("( 15:04 ) ")
				FormatedTime += "-|- Arrived.\n "
				entry.Text += FormatedTime
				entry.MultiLine = true
				entry.Refresh()

				// GORUTINE
				go liveClock(c, timeDoneText, quit, endTime, arrivedTime, timeLeftText, latestArchivedTime)
				liveClockActive = true

				// WRITE TO LOG
				writeToLog(FormatedTime)
			}

		}),
		widget.NewToolbarAction(theme.MediaRecordIcon(), func() {
			if arrActive && !eodActive {
				if selectEntry.Text != "" {
					currentTime := time.Now()
					if firstTimeArchivedTime {
						latestArchivedTimeZoneFix := latestArchivedTime.Add(-1 * time.Hour)
						timeWork += currentTime.Sub(latestArchivedTimeZoneFix)
						firstTimeArchivedTime = false
					} else {
						timeWork += currentTime.Sub(latestArchivedTime)
					}

					fmt.Println(latestArchivedTime, currentTime)

					TEMP := fmt.Sprintf("%d h %d min", int(timeWork.Hours()), int(timeWork.Minutes())%60)
					timeWorkText.Text = "Work Done: " + TEMP
					timeWorkText.Refresh()

					entry.TextStyle = fyne.TextStyle{Bold: true}
					FormatedTime := latestArchivedTime.Format("( 15:04 - ")
					FormatedTime += currentTime.Format("15:04 )")

					FormatedTime += " -|- Task: "
					FormatedTime += selectEntry.Text
					FormatedTime += " -|- Summary: "
					FormatedTime += summaryEntry.Text + " \n"
					entry.Text += FormatedTime

					latestArchivedTime = currentTime
					entry.MultiLine = true
					selectEntry.Text = ""
					selectEntry.FocusLost()

					summaryEntry.Text = ""
					summaryEntry.FocusLost()

					entry.Refresh()

					log.Println("Task:", selectEntry, FormatedTime)
					writeToLog(FormatedTime)

				}
			}
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if arrActive && !eodActive {
				currentTime := time.Now()
				entry.TextStyle = fyne.TextStyle{Bold: true}
				FormatedTime := latestArchivedTime.Format("( 15:04 -")
				FormatedTime += currentTime.Format(" 15:04 )")
				FormatedTime += " -|- Break: \n"
				entry.Text += FormatedTime

				if firstTimeArchivedTime {
					latestArchivedTimeZoneFix := latestArchivedTime.Add(-1 * time.Hour)
					timeSlacking += currentTime.Sub(latestArchivedTimeZoneFix)
					firstTimeArchivedTime = false
				} else {
					timeSlacking += currentTime.Sub(latestArchivedTime)
				}

				TEMP := fmt.Sprintf("%d h %d min", int(timeSlacking.Hours()), int(timeSlacking.Minutes())%60)
				TEMP2 := removeChar(TEMP, '-')
				timeStatText.Text = "Slacking: " + TEMP2
				timeStatText.Refresh()

				latestArchivedTime = currentTime
				entry.MultiLine = true
				entry.Refresh()
				log.Println("Break:", entry.Text, FormatedTime)

				// Write to log
				writeToLog(FormatedTime)
			}

		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			if !eodActive && arrActive {
				eodActive = true
				currentTime := time.Now()
				timeDone = arrivedTime.Sub(currentTime)
				entry.TextStyle = fyne.TextStyle{Bold: true}
				FormatedTime := arrivedTime.Format("( 15:04 -")
				FormatedTime += currentTime.Format(" 15:04 )")
				FormatedTime += " -|- End of Day. \n"
				entry.Text += FormatedTime

				timeDone = latestArchivedTime.Sub(currentTime)
				TEMP := fmt.Sprintf("%d h %d min", int(timeDone.Hours()), int(timeDone.Minutes())%60)
				TEMP2 := removeChar(TEMP, '-')
				timeDoneText.Text = "Worked Today: " + TEMP2
				if liveClockActive {
					quit <- true
				}
				if liveClockOverrideActive {
					quit0 <- true
				}
				timeDoneText.Refresh()

				entry.MultiLine = true
				entry.Refresh()

				// write to log
				writeToLog(FormatedTime)

				writeToLog(timeWorkText.Text)
				writeToLog(timeStatText.Text)
				writeToLog(timeDoneText.Text)

			}

		}),
	)

	// OPTIONS MENU --------------------------------------
	endTimeOP := widget.NewEntry()
	TEMP5 := endTime.Format("15:04")
	endTimeOP.Text = TEMP5
	timeOptionLabel := widget.NewLabel("When do you end work?")
	// WHAT TIME DO YOU END? OPTIONS
	startTimeOP := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		TEMP := endTime.Add(+30 * time.Minute)
		endTime = TEMP
		TEMP0 := endTime.Format("15:04")
		endTimeOP.Text = TEMP0
		endTimeOP.Refresh()
	})
	timeMinusOP := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {
		TEMP1 := endTime.Add(-30 * time.Minute)
		endTime = TEMP1
		TEMP2 := endTime.Format("15:04")
		endTimeOP.Text = TEMP2
		endTimeOP.Refresh()
	})
	// OPTION SAVE BUTTON
	saveOptions := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		// ENDTIME OPTION && API KEY
		if apiKeyEntry.Text != "" {
			apiKey = apiKeyEntry.Text
			apiKey += "\n"
		}

		if projectIDEntry.Text != "" {
			projectID = projectIDEntry.Text
		}

		TEMP4 := endTime.Format("15:04\n")

		data := TEMP4 + apiKey + projectID
		newData := []byte(data)
		err := os.WriteFile("useroptions.txt", newData, 0644)
		if err != nil {
			fmt.Println("Error while saving options")
		}
		// SAVE STATUS TEXT
		projectIDLabel.Text = "SAVED"
		ticker := time.NewTicker(3 * time.Second)
		<-ticker.C
		projectIDLabel.Text = ""
	})

	// MARKED DATES CHECK

	FormatedYCT10 := yct.Format("(2006-01-02)")
	clearDate := fileContents("markedDates.txt", FormatedYCT10)
	if clearDate {
		restrictMarkIcon.Show()
		markButton.Disable()
	} else {
		restrictMarkIcon.Hide()
		markButton.Enable()
	}

	// HIDE OTHER MENUS FROM START

	endTimeOP.Hide()
	timeOptionLabel.Hide()
	startTimeOP.Hide()
	timeMinusOP.Hide()
	saveOptions.Hide()

	apiKeyLabel.Hide()
	apiKeyEntry.Hide()

	//projectIDLabel.Hide()
	projectIDEntry.Hide()

	sendToGitButton.Hide()
	sendToGitEntry.Hide()

	selectEntry.PlaceHolder = "Select Task"
	selectEntry.Refresh()

	summaryEntry.PlaceHolder = "Summary"
	summaryEntry.Refresh()

	// MENU SELECT
	toolbarSide := widget.NewToolbar(
		// LOG MENU
		widget.NewToolbarAction(theme.GridIcon(), func() {
			sidebar.Show()
			selectEntry.Show()
			summaryEntry.Show()

			endTimeOP.Hide()
			timeOptionLabel.Hide()
			startTimeOP.Hide()
			timeMinusOP.Hide()
			saveOptions.Hide()

			apiKeyLabel.Hide()
			apiKeyEntry.Hide()

			//projectIDLabel.Hide()
			projectIDEntry.Hide()

			sendToGitButton.Hide()
			sendToGitEntry.Hide()
		}),
		// SEND MENU
		widget.NewToolbarAction(theme.MailSendIcon(), func() {
			sidebar.Hide()
			selectEntry.Hide()
			summaryEntry.Hide()

			endTimeOP.Hide()
			timeOptionLabel.Hide()
			startTimeOP.Hide()
			timeMinusOP.Hide()
			saveOptions.Hide()

			apiKeyLabel.Hide()
			apiKeyEntry.Hide()

			//projectIDLabel.Hide()
			projectIDEntry.Hide()

			sendToGitButton.Show()
			sendToGitEntry.Show()

			loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
		}),
		// OPTIONS MENU
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			sidebar.Hide()
			selectEntry.Hide()
			summaryEntry.Hide()

			endTimeOP.Show()
			timeOptionLabel.Show()
			startTimeOP.Show()
			timeMinusOP.Show()
			saveOptions.Show()

			apiKeyLabel.Show()
			apiKeyEntry.Show()

			//projectIDLabel.Show()
			projectIDEntry.Show()

			sendToGitButton.Hide()
			sendToGitEntry.Hide()
		}),
	)

	// ENTRY LOG ---------------------------------------------------------------------------- LOAD IN LOG TO ENTRY
	tn := time.Now()
	tnf := tn.Format("(2006-01-02)")
	var workTime time.Duration
	var totalWorkTime time.Duration
	var slackingTime time.Duration
	var totalSlackingTime time.Duration

	logFile, err := os.Open("log.txt")
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
			// fix latest time
			latestArchivedTime = updateLatestArchivedTime(line, latestArchivedTime)

			// status text
			shiftLeft = updateShiftLeftText(endTime)
			slackingTime += updateSlackingStatus(line, totalSlackingTime)
			workTime += updateWorkDoneStatus(line, totalWorkTime)
			eod := strings.Contains(line, "End of Day")
			if eod {
				eodActive = true
			}
			arr := strings.Contains(line, "Arrived")
			if arr {
				arrActive = true
			}

			if !liveClockActive && !liveClockOverrideActive {
				go LiveClockOverride(ch, quit0, line, shiftLeft, endTime, timeLeftText, arrivedTime, timeDoneText)
				liveClockOverrideActive = true
				fmt.Println(liveClockOverrideActive)
			}
			// Write to Entry
			entry.MultiLine = true
			entry.Text += new + "\n"
			entry.Refresh()
		}

	}
	timeWork = workTime
	timeSlacking = slackingTime

	TEMP := "Work Done: "
	TEMP += fmt.Sprintf("%d h %d min", int(workTime.Hours()), int(workTime.Minutes())%60)
	timeWorkText.Text = TEMP
	timeWorkText.Refresh()

	TEMP12 := "Slacking: "
	TEMP12 += fmt.Sprintf("%d h %d min", int(slackingTime.Hours()), int(slackingTime.Minutes())%60)
	timeStatText.Text = TEMP12
	timeStatText.Refresh()

	TEMP13 := "Shift Left: "
	TEMP13 += fmt.Sprintf("%d h %d min ", int(shiftLeft.Hours()), int(shiftLeft.Minutes())%60)
	timeLeftText.Text = TEMP13
	timeLeftText.Refresh()

	// MAIN ENTRY SAVE BUTTON
	image := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		gg := strings.Contains(entry.Text, "\n")
		fmt.Println(gg)
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
				fmt.Printf("( was located at position %d", position)
				value = value[position:]
				value = value[:]
				hf = append(hf, value)
				fmt.Println(value)
			} else {
				fmt.Println("( was not located")
			}

		}
		var hn []string
		FYCT := yct.Format("(2006-01-02)")
		for _, v := range hf {
			hn = append(hn, FYCT+" "+v)
		}
		fmt.Println(hn)
		// Remove date from log
		removeLineFromFile("log.txt", FYCT)
		// Write new date infro into log
		for _, v := range hn {
			newWriteToLog(v)
		}
		entry.Text += "\n"
		// SAVE STATUS TEXT
		projectIDLabel.Text = "SAVED"
		ticker := time.NewTicker(3 * time.Second)
		<-ticker.C
		projectIDLabel.Text = ""
	})

	// SEND UI ELEMENTS TO GET POSITIONAL AND SIZE DATA TO layout.go

	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}
	obj := []fyne.CanvasObject{content, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, toolbar, sidebar, entry, selectEntry, summaryEntry, dividers[0], dividers[1], dividers[2], image}
	return container.New(nykonstruktorLayout(toolbar, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, summaryEntry, content, image, dividers), obj...)
}
