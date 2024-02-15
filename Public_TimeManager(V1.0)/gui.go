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

	// SAVES ------------------------------------------------------------
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
		//fmt.Println(line)
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
		//fmt.Println(i)
		//fmt.Println(apiKey)

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

	// UI -----------------------------------------------------------------------

	// DECLARITIONER AV VARIABLER
	//right := widget.NewLabel("Right")
	content := canvas.NewRectangle(color.Gray{Y: 0xee})
	//image.FillMode = canvas.ImageFillContain
	entry := widget.NewMultiLineEntry()
	selectEntryList := []string{"Enter a valid API key in options"}
	selectEntry := widget.NewSelectEntry(selectEntryList)
	//UpdateSelectEntryOptions(selectEntry)
	readyToSendLog := []string{}
	restrictMarkIcon := widget.NewIcon(theme.CancelIcon())

	sendToGitEntry := widget.NewMultiLineEntry()
	/*
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
					task = task[:]
					fmt.Println(task)
					SendToGit(apiKey, projectID, duration, getIIDFromTask(task))
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

		})
	*/

	apiRequest(selectEntry, apiKey, projectID)
	apiKeyLabel := widget.NewLabel("Enter your gitlab API key here")
	apiKeyEntry := widget.NewPasswordEntry()

	projectIDLabel := widget.NewLabel("Enter your gitlab Project ID here")
	projectIDEntry := widget.NewEntry()

	date := widget.NewLabel("")
	ct := time.Now()
	date.Text = ct.Weekday().String()
	date.Text += ct.Format(", 2006-01-02 ")

	yct := ct
	c := make(chan int)
	quit := make(chan bool)

	ch := make(chan int)
	quit0 := make(chan bool)

	var liveClockActive bool
	var liveClockOverrideActive bool

	var arrivedTime time.Time
	var latestArchivedTime time.Time
	var firstTimeArchivedTime bool = true

	// One per day restriction section
	var eodActive bool = false
	var arrActive bool = false

	//var timeLeft time.Duration
	var timeSlacking time.Duration
	var timeWork time.Duration
	var timeDone time.Duration
	var shiftLeft time.Duration

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

	// DATE NAVIGATION ----------------------------------------------------------------------------DATE NAVIGATION

	markButton := widget.NewCheck("", func(b bool) {
		FormatedYCT10 := yct.Format("(2006-01-02)")
		clearDate := fileContents("markedDates.txt", FormatedYCT10)
		//fmt.Println(FormatedYCT10)
		if !clearDate {
			if b {
				//fmt.Println("Checked")
				FormatedYCT := yct.Format("(2006-01-02)")
				//fmt.Println(FormatedYCT)
				readyToSendLog = append(readyToSendLog, FormatedYCT)
				//fmt.Println(readyToSendLog)
			} else if !b {
				//fmt.Println("Unchecked")
				FormatedYCT0 := yct.Format("(2006-01-02)")
				//fmt.Println(FormatedYCT0)
				readyToSendLog = removeElementFromList(readyToSendLog, FormatedYCT0)
				readyToSendLog = removeDuplicates(readyToSendLog)
				//fmt.Println(readyToSendLog)
			}
		}
	})

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
				task = task[:]
				fmt.Println(task)
				SendToGit(apiKey, projectID, duration, getIIDFromTask(task))
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
				//log.Println("Arrived:", entry.Text, FormatedTime)
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
					//fmt.Println(timeWork)

					//fmt.Println(timeWork)
					TEMP := fmt.Sprintf("%d h %d min", int(timeWork.Hours()), int(timeWork.Minutes())%60)
					//TEMP2 := removeChar(TEMP, '-')
					timeWorkText.Text = "Work Done: " + TEMP
					timeWorkText.Refresh()

					entry.TextStyle = fyne.TextStyle{Bold: true}
					FormatedTime := latestArchivedTime.Format("( 15:04 - ")
					FormatedTime += currentTime.Format("15:04 )")

					FormatedTime += " -|- Task: "
					FormatedTime += selectEntry.Text + " \n"
					entry.Text += FormatedTime

					latestArchivedTime = currentTime
					entry.MultiLine = true
					selectEntry.Text = ""
					selectEntry.FocusLost()
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

				//fmt.Println(latestArchivedTime, currentTime)
				//latestArchivedTimeCET := latestArchivedTime.Add(-1 * time.Hour)
				//timeSlacking += latestArchivedTime.Sub(currentTime)
				//fmt.Println(timeSlacking)

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
				//quit <- true
				//quit0 <- true
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

	// Options Menu --------------------------------------
	endTimeOP := widget.NewEntry()
	TEMP5 := endTime.Format("15:04")
	endTimeOP.Text = TEMP5

	//var TEMP3 time.Time
	//endTime = TEMP3.Add(16 * time.Hour)
	timeOptionLabel := widget.NewLabel("When do you end work?")

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

	saveOptions := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		// ENDTIME OPTION && API KEY
		if apiKeyEntry.Text != "" {
			apiKey = apiKeyEntry.Text
			apiKey += "\n"
		}
		//data0 := []byte(apiKey)

		if projectIDEntry.Text != "" {
			projectID = projectIDEntry.Text
		}

		TEMP4 := endTime.Format("15:04\n")
		//data := []byte(TEMP4)

		data := TEMP4 + apiKey + projectID
		newData := []byte(data)
		err := os.WriteFile("useroptions.txt", newData, 0644)
		if err != nil {
			fmt.Println("Error while saving options")
		}
	})

	FormatedYCT10 := yct.Format("(2006-01-02)")
	clearDate := fileContents("markedDates.txt", FormatedYCT10)
	if clearDate {
		restrictMarkIcon.Show()
		markButton.Disable()
	} else {
		restrictMarkIcon.Hide()
		markButton.Enable()
	}

	endTimeOP.Hide()
	timeOptionLabel.Hide()
	startTimeOP.Hide()
	timeMinusOP.Hide()
	saveOptions.Hide()

	apiKeyLabel.Hide()
	apiKeyEntry.Hide()

	projectIDLabel.Hide()
	projectIDEntry.Hide()

	sendToGitButton.Hide()
	sendToGitEntry.Hide()

	//restrictMarkIcon.Hide()
	//seperator := widget.NewSeparator()

	toolbarSide := widget.NewToolbar(

		widget.NewToolbarAction(theme.GridIcon(), func() {
			sidebar.Show()
			selectEntry.Show()

			endTimeOP.Hide()
			timeOptionLabel.Hide()
			startTimeOP.Hide()
			timeMinusOP.Hide()
			saveOptions.Hide()

			apiKeyLabel.Hide()
			apiKeyEntry.Hide()

			projectIDLabel.Hide()
			projectIDEntry.Hide()

			sendToGitButton.Hide()
			sendToGitEntry.Hide()
		}),

		widget.NewToolbarAction(theme.MailSendIcon(), func() {
			sidebar.Hide()
			selectEntry.Hide()

			endTimeOP.Hide()
			timeOptionLabel.Hide()
			startTimeOP.Hide()
			timeMinusOP.Hide()
			saveOptions.Hide()

			apiKeyLabel.Hide()
			apiKeyEntry.Hide()

			projectIDLabel.Hide()
			projectIDEntry.Hide()

			sendToGitButton.Show()
			sendToGitEntry.Show()

			loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
		}),

		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			sidebar.Hide()
			selectEntry.Hide()

			endTimeOP.Show()
			timeOptionLabel.Show()
			startTimeOP.Show()
			timeMinusOP.Show()
			saveOptions.Show()

			apiKeyLabel.Show()
			apiKeyEntry.Show()

			projectIDLabel.Show()
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
			//fmt.Println(tnf)
			//fmt.Println(new)
			// fix arrived time
			arrivedTime = updateArrivedTime(line, arrivedTime)
			// fix latest time
			latestArchivedTime = updateLatestArchivedTime(line, latestArchivedTime)

			// status text
			shiftLeft = updateShiftLeftText(endTime)
			slackingTime += updateSlackingStatus(line, totalSlackingTime)
			//updateWorkedTodayStatus(arrivedTime)
			workTime += updateWorkDoneStatus(line, totalWorkTime)
			//go liveClock(c, timeDoneText, quit, endTime, arrivedTime, timeLeftText, latestArchivedTime)
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

			//fmt.Println(liveClockOverrideActive)
			// Write to Entry
			entry.MultiLine = true
			entry.Text += new + "\n"
			entry.Refresh()
		}

	}
	timeWork = workTime
	//fmt.Println(timeWork)
	//fmt.Println(workTime)
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

	/*
		fmt.Println(TEMP)
		fmt.Print("SlackingTime: ")
		fmt.Println(slackingTime)
		fmt.Print("LatestArchivedTime: ")
		fmt.Println(latestArchivedTime)
		fmt.Print("arrived: ")
		fmt.Println(arrivedTime)
	*/

	image := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		//Form entry.text adn add date(yct)

		//entry.Text = entry.Text[:len(entry.Text)-1]
		//lastChar := entry.Text[len(entry.Text)-1:]
		//fmt.Println("LASTCHAR:" + lastChar)
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

		// Remove : and spaces

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

		//fmt.Println(hh)

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
	})

	//left.Text = "HEllo"

	//content.Alignment = fyne.TextAlignCenter

	//return container.NewBorder(toolbar, nil, left, right, content)

	dividers := [3]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
	}
	obj := []fyne.CanvasObject{content, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, toolbar, sidebar, entry, selectEntry, dividers[0], dividers[1], dividers[2], image}
	return container.New(nykonstruktorLayout(toolbar, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, content, image, dividers), obj...)
}

func removeChar(inputString string, deleteChar rune) string {
	result := ""
	for _, char := range inputString {
		if char != deleteChar {
			result += string(char)
		}
	}
	return result
}

func liveClock(c chan int, a *canvas.Text, quit chan bool, endTime time.Time, arrivedTime time.Time, timeLeftText *canvas.Text, latestAchivedTime time.Time) {
	ticker := time.NewTicker(1 * time.Second)
	var timer time.Time
	TEMP4 := arrivedTime.Format("15:04")
	TEMP8, err := time.Parse("15:04", TEMP4)

	if err != nil {
		fmt.Println("Error trying to parse TEMP4 to time.Time")
	}
	for {
		select {
		case <-quit:
			return
		default:
			// WORKED TODAY
			<-ticker.C
			timer = timer.Add(+1 * time.Second)
			TEMP := timer.Format(" 15:04")
			Dis := "Worked Today:"
			Dis += TEMP
			a.Text = Dis
			a.Refresh()
			//fmt.Println(timer)
			// SHIFT LEFT
			TEMP10 := TEMP8
			TEMP8 = TEMP10.Add(+1 * time.Second)
			timeLeft := endTime.Sub(TEMP8)

			TEMP6 := fmt.Sprintf("%d h %d min ", int(timeLeft.Hours()), int(timeLeft.Minutes())%60)
			TEMP7 := "Shift Left: "
			TEMP7 += TEMP6
			timeLeftText.Text = TEMP7
			timeLeftText.Refresh()
		}
	}
}

func LiveClockOverride(ch chan int, quit chan bool, line string, shiftLeft time.Duration, endTime time.Time, timeLeftText *canvas.Text, arrivedTime time.Time, timeDoneText *canvas.Text) {
	tf := strings.Contains(line, "Arrived")
	ticker := time.NewTicker(1 * time.Second)
	// setting endtime up to date(UTD)
	currentTime := time.Now()
	clone := endTime
	endTime = clone.Add(-1 * time.Hour)
	endFormated := endTime.Format("15:04")
	currentFormated := currentTime.Format("2006-01-02 ")
	currentFormated += endFormated
	endTimeUTD, err := time.Parse("2006-01-02 15:04", currentFormated)
	// Fixing arrived time
	clone0 := arrivedTime
	arrivedTime = clone0.Add(-1 * time.Hour)
	if err != nil {
		fmt.Println("Error while trying to parse currentFormated in LiveClockOverride")
	}
	// RUNNING CLOCK
	if tf {
		for {
			select {
			case <-quit:
				return
			default:
				<-ticker.C
				// SHIFT LEFT
				// LOGIC
				currentTime = time.Now()
				ShiftLeft := endTimeUTD.Sub(currentTime)
				// STATUS UPDATE
				TEMP := fmt.Sprintf("%d h %d min ", int(ShiftLeft.Hours()), int(ShiftLeft.Minutes())%60)
				TEMP0 := "Shift Left: "
				TEMP0 += TEMP
				timeLeftText.Text = TEMP0
				timeLeftText.Refresh()

				// WORKED DONE
				WorkDone := currentTime.Sub(arrivedTime)

				TEMP2 := fmt.Sprintf("%d h %d min ", int(WorkDone.Hours()), int(WorkDone.Minutes())%60)
				TEMP3 := "Worked Today: "
				TEMP3 += TEMP2
				timeDoneText.Text = TEMP3
				timeDoneText.Refresh()
			}
		}
	} else {
		fmt.Println("NO Arrived exists in log")
	}

}

func writeToLog(FormatedTime string) {
	var TEMP string
	logFile, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		TEMP += line + "\n"
	}
	gg := time.Now()
	gt := gg.Format("(2006-01-02) ")
	gt += FormatedTime
	TEMP += gt
	data := []byte(TEMP)
	gerr := os.WriteFile("log.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while saving Arrived to log")
	}
}

func readFromLog(date time.Time, entry *widget.Entry) {
	entry.Text = ""
	tnf := date.Format("(2006-01-02)")

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
			entry.MultiLine = true
			entry.Text += ": " + new + "\n"
			entry.Refresh()
		}
		entry.Refresh()
	}
}

func takeTextFromString(text string, start, end int) string {
	return strings.TrimSpace(text[start:end])
}

func updateShiftLeftText(endTime time.Time) time.Duration {
	// SHIFT LEFT
	currentTime := time.Now()
	TEMP := endTime.Format("15:04")
	TEMP0 := currentTime.Format("2006-01-02")
	TEMP0 += " " + TEMP
	newEndTime, err := time.Parse("2006-01-02 15:04", TEMP0)
	if err != nil {
		fmt.Println("Error while trying to parse endTime in function updateStatusText")
	}
	shiftLeft := newEndTime.Sub(currentTime)
	//SLtext := fmt.Sprintf("%d h %d min ", int(shiftLeft.Hours()), int(shiftLeft.Minutes())%60)
	//fmt.Println("SHIFT LEFT:" + SLtext)
	return shiftLeft
}

func updateSlackingStatus(line string, totalSlackingTime time.Duration) time.Duration {
	currentTime := time.Now()
	TF0 := strings.Contains(line, "Break")
	if TF0 {
		FcurrentTime := currentTime.Format("2006-01-02")
		N10 := takeTextFromString(line, 15, 21)
		N20 := takeTextFromString(line, 23, 29)
		BN10 := FcurrentTime + " " + N10
		BN20 := FcurrentTime + " " + N20
		fmt.Println(BN20)
		NN10, errr0 := time.Parse("2006-01-02 15:04", BN10)
		if errr0 != nil {
			fmt.Println("Error while trying to parse N1")
		}
		NN20, errrr := time.Parse("2006-01-02 15:04", BN20)
		if errrr != nil {
			fmt.Println("Error while trying to parse N2")
		}
		totalSlackingTime = NN20.Sub(NN10)
	}
	return totalSlackingTime
}
func updateWorkedTodayStatus(arrivedTime time.Time) {
	currentTime := time.Now()
	workedToday := currentTime.Sub(arrivedTime)
	WTtext := fmt.Sprintf("%d h %d min ", int(workedToday.Hours()), int(workedToday.Minutes())%60)
	fmt.Println("WORKED TODAY: " + WTtext)
}
func updateWorkDoneStatus(line string, totalWorkTime time.Duration) time.Duration {
	currentTime := time.Now()
	TF := strings.Contains(line, "Task")
	if TF {
		FcurrentTime := currentTime.Format("2006-01-02")
		N1 := takeTextFromString(line, 15, 20)
		N2 := takeTextFromString(line, 22, 28)
		BN1 := FcurrentTime + " " + N1
		BN2 := FcurrentTime + " " + N2
		NN1, errr := time.Parse("2006-01-02 15:04", BN1)
		if errr != nil {
			fmt.Println("Error while trying to parse N1")
		}
		NN2, errrr := time.Parse("2006-01-02 15:04", BN2)
		if errrr != nil {
			fmt.Println("Error while trying to parse N2")
		}
		totalWorkTime = NN2.Sub(NN1)
	}
	return totalWorkTime
}

func updateLatestArchivedTime(line string, latestArchivedTime time.Time) time.Time {
	TFA := strings.Contains(line, "Arrived")
	if TFA {
		FLatestTimeHM := takeTextFromString(line, 14, 21)
		//fmt.Println(FLatestTimeHM)
		FLatestTimeYMD := takeTextFromString(line, 1, 11)
		//fmt.Println(FLatestTimeYMD)
		FLatestTimeYMD += " " + FLatestTimeHM
		time0, err := time.Parse("2006-01-02 15:04", FLatestTimeYMD)
		if err != nil {
			fmt.Println("Error while parsing line to Latest time")
		}
		latestArchivedTime = time0
	} else {
		FLatestTimeHM := takeTextFromString(line, 23, 28)
		//fmt.Println(FLatestTimeHM)
		FLatestTimeYMD := takeTextFromString(line, 1, 11)
		//fmt.Println(FLatestTimeYMD)
		FLatestTimeYMD += " " + FLatestTimeHM
		time0, err := time.Parse("2006-01-02 15:04", FLatestTimeYMD)
		if err != nil {
			fmt.Println("Error while parsing line to Latest time")
		}

		latestArchivedTime = time0
	}

	return latestArchivedTime
}

func updateArrivedTime(line string, arrivedTime time.Time) time.Time {
	arrived := strings.Contains(line, "Arrived")
	if arrived {
		FarrivedTimeHM := takeTextFromString(line, 14, 21)
		//fmt.Println("HM: " + FarrivedTimeHM)
		FarrivedTimeYMD := takeTextFromString(line, 1, 11)
		//fmt.Println("YMD: " + FarrivedTimeYMD)
		FarrivedTimeYMD += " " + FarrivedTimeHM
		//fmt.Println("YMD-HM: " + FarrivedTimeYMD)
		time, err := time.Parse("2006-01-02 15:04", FarrivedTimeYMD)
		if err != nil {
			fmt.Println("Error while parsing arrived time")
		}
		//fmt.Println(time)
		arrivedTime = time
		//go liveClock(c, timeDoneText, quit, endTime, arrivedTime, timeLeftText, latestArchivedTime)

	}
	return arrivedTime
}

func UpdateSelectEntryOptions(selectEntry *widget.SelectEntry) {
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
		fmt.Println(selectEntryList)
	}
}

func removeElementFromList(list []string, v string) []string {
	for index, str := range list {
		if str == v {
			list = append(list[:index], list[index+1:]...)
			break
		}
	}
	return list
}

func removeDuplicates(list []string) []string {
	Map := make(map[string]bool)

	tempList := []string{}

	for _, v := range list {
		if _, ok := Map[v]; !ok {
			tempList = append(tempList, v)
			Map[v] = true
		}
	}
	return tempList
}

func loadReadyDatesToEntry(readyToSendLog []string, sendToGitEntry *widget.Entry) {
	sendToGitEntry.Text = ""
	sendToGitEntry.Text = "<> Press upload button to log time to gitlab <>"
	for num, v := range readyToSendLog {
		fmt.Println(v, " ")
		fmt.Print(num)
		v0 := v
		v += "\n"
		sendToGitEntry.Text += "\n"
		sendToGitEntry.Text += v

		File, err := os.Open("log.txt")
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(File)
		TEMP9 := ""
		data := []byte(TEMP9)
		gerr := os.WriteFile("response.txt", data, 0644)
		if gerr != nil {
			fmt.Println("Error while writeing to Tasks")
		}

		for scanner.Scan() {
			line := scanner.Text()
			Clear := strings.Contains(line, v0)
			if Clear {
				new := strings.Trim(line, v0)

				taskTrue := strings.Contains(new, "Task")
				if taskTrue {
					firstPreFormatesTaskTime := takeTextFromString(new, 3, 8)
					secondPreFormatesTaskTime := takeTextFromString(new, 11, 16)
					fmt.Println(firstPreFormatesTaskTime)
					fmt.Println(secondPreFormatesTaskTime)
					first, err := time.Parse("15:04", firstPreFormatesTaskTime)
					if err != nil {
						fmt.Println("Error while trying to parse in loadReadyDatesToEntry")
					}
					second, err0 := time.Parse("15:04", secondPreFormatesTaskTime)
					if err0 != nil {
						fmt.Println("Error while trying to parse in loadReadyDatesToEntry")
					}
					fmt.Println(first)
					fmt.Println(second)
					taskTime := second.Sub(first)
					fmt.Println(taskTime)
					formatedTaskTime := fmt.Sprintf("%dh%dm", int(taskTime.Hours()), int(taskTime.Minutes())%60)
					fmt.Println(formatedTaskTime)
					fmt.Println(new)
					tempdelText := takeTextFromString(new, 0, 28)
					delText := strings.Trim(new, tempdelText)
					fmt.Println(tempdelText)
					fmt.Println(delText)
					formatedTaskTime += " --> Task: " + delText

					sendToGitEntry.MultiLine = true
					sendToGitEntry.Text += formatedTaskTime + "\n"
					writeToTasks(formatedTaskTime+"\n", "response.txt")
					sendToGitEntry.Refresh()
				}
			}
		}

		sendToGitEntry.Refresh()
	}
}

func fileContents(file string, content string) bool {
	File, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var clear bool
	scanner := bufio.NewScanner(File)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line + "  :: " + content)
		clear = strings.Contains(line, content)
		if clear {
			return clear
		}
	}
	return clear
}

func removeLineFromFile(file string, lineToRemove string) {
	var logLineList []string
	File, err := os.Open(file)
	if err != nil {
		fmt.Println("Error while trying to remove string from file")
	}

	scanner := bufio.NewScanner(File)
	for scanner.Scan() {
		line := scanner.Text()

		gfg := strings.Contains(line, lineToRemove)
		if gfg {
			//continue
		} else {
			logLineList = append(logLineList, line)
		}

	}
	// Clear log
	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile(file, data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	for _, v := range logLineList {
		writeToTasks(v, file)
	}
}

func newWriteToLog(Line string) {
	var TEMP string
	logFile, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		TEMP += line + "\n"
	}
	TEMP += Line
	data := []byte(TEMP)
	gerr := os.WriteFile("log.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while saving Arrived to log")
	}
}
