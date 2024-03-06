package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func FunctionHolder() {

}

// TIME HOLDER FUNCTIONS

func liveClock(c chan int, a *canvas.Text, quit chan bool, endTime time.Time, arrivedTime time.Time, timeLeftText *canvas.Text, latestAchivedTime time.Time) {
	ticker := time.NewTicker(1 * time.Second)
	var timer time.Time
	TEMP4 := arrivedTime.Format("15:04")
	TEMP8, err := time.Parse("15:04", TEMP4)

	var RFT time.Duration
	var TEMP67 time.Time
	var DurationHolder time.Duration

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
			TEMP67 = timer
			timer = timer.Add(+1 * time.Second)

			RFT = timer.Sub(TEMP67)
			DurationHolder += RFT
			TEMP68 := fmt.Sprintf("%d h %d min ", int(DurationHolder.Hours()), int(DurationHolder.Minutes())%60)
			Dis := "Worked Today: "
			Dis += TEMP68
			a.Text = Dis
			a.Refresh()
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

// FILE HANDELING FUNCTIONS

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
func fileContents(file string, content string) bool {
	File, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var clear bool
	scanner := bufio.NewScanner(File)
	for scanner.Scan() {
		line := scanner.Text()
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

// LOGICAL FUNCTIONS/ HELPER FUNCTIONS

func removeChar(inputString string, deleteChar rune) string {
	result := ""
	for _, char := range inputString {
		if char != deleteChar {
			result += string(char)
		}
	}
	return result
}
func takeTextFromString(text string, start, end int) string {
	return strings.TrimSpace(text[start:end])
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
func GetSummaryFromTask(task string) string {
	//var sumHolder string
	File, err := os.Open("response.txt")
	if err != nil {
		panic(err)
	}
	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()
		clear := strings.Contains(line, task)
		if clear {
			clear0 := strings.Contains(line, "Summary")
			if clear0 {
				line = line[len(task):]
				line = line[14:]
				line = strings.Replace(line, " ", "+", -1)
				return line
			}
		}
	}
	return "Error"
}

// STATUS TEXT UPDATE FUNCTIONS

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
		FLatestTimeYMD := takeTextFromString(line, 1, 11)
		FLatestTimeYMD += " " + FLatestTimeHM
		time0, err := time.Parse("2006-01-02 15:04", FLatestTimeYMD)
		if err != nil {
			fmt.Println("Error while parsing line to Latest time")
		}
		latestArchivedTime = time0
	} else {
		FLatestTimeHM := takeTextFromString(line, 23, 28)
		FLatestTimeYMD := takeTextFromString(line, 1, 11)
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
		FarrivedTimeYMD := takeTextFromString(line, 1, 11)
		FarrivedTimeYMD += " " + FarrivedTimeHM
		time, err := time.Parse("2006-01-02 15:04", FarrivedTimeYMD)
		if err != nil {
			fmt.Println("Error while parsing arrived time")
		}
		arrivedTime = time
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
func loadReadyDatesToEntry(readyToSendLog []string, sendToGitEntry *widget.Entry) {
	sendToGitEntry.Text = ""
	for _, v := range readyToSendLog {
		sendToGitEntry.Text = "<> Press upload button to log time to gitlab <>"
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
					first, err := time.Parse("15:04", firstPreFormatesTaskTime)
					if err != nil {
						fmt.Println("Error while trying to parse in loadReadyDatesToEntry")
					}
					second, err0 := time.Parse("15:04", secondPreFormatesTaskTime)
					if err0 != nil {
						fmt.Println("Error while trying to parse in loadReadyDatesToEntry")
					}
					taskTime := second.Sub(first)
					formatedTaskTime := fmt.Sprintf("%dh%dm", int(taskTime.Hours()), int(taskTime.Minutes())%60)
					tempdelText := takeTextFromString(new, 0, 28)
					delText := strings.Replace(new, tempdelText, "", -1)

					formatedTaskTime += " --> Task:" + delText

					var summary string
					// Seperate summary from Task Name
					if strings.Contains(formatedTaskTime, "-|-") {
						positionOfChar := strings.IndexRune(formatedTaskTime, '|')
						summary = takeTextFromString(formatedTaskTime, positionOfChar-1, len(formatedTaskTime))
						formatedTaskTime = strings.Replace(formatedTaskTime, summary, "", -1)
					}
					// make delText == task name
					delText = strings.Replace(delText, summary, "", -1)
					delText = delText[2:]

					sendToGitEntry.MultiLine = true
					sendToGitEntry.Text += formatedTaskTime + "\n"
					sendToGitEntry.Text += summary + "\n" + "\n"
					writeToTasks(formatedTaskTime+"\n", "response.txt")
					writeToTasks(delText+summary+"\n", "response.txt")
					sendToGitEntry.Refresh()
				}
			}
		}

		sendToGitEntry.Refresh()
	}
}
