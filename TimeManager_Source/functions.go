package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// ProgressBar
var NumofTasks int
var TaskTakeValue float64

// helper text
var HelperTextActive bool
var HelperTextDuration int
var LoaderActive bool

// BACKGROUND CHECKS
func SystemCheck(c chan int, exit chan bool, PlayButton *widget.Button, RecordButton *widget.Button) {
	ticker := time.NewTicker(1 * time.Second)
	var BackBackColor bool
	for {
		select {
		case <-exit:
			return
		default:
			// WORKED TODAY
			<-ticker.C
			if checkArrivedStatus() {
				PlayButton.Disable()
				RecordButton.Enable()
			} else {
				PlayButton.Enable()
				arrActive = false
				RecordButton.Disable()
			}

			if !BackBackColor {
				popBackBackground.FillColor = color.RGBA{204, 204, 204, 255}
				BackBackColor = true
				popBackBackground.Refresh()
			} else {
				popBackBackground.FillColor = color.RGBA{150, 150, 150, 255}
				BackBackColor = false
				popBackBackground.Refresh()
			}

			if HelperTextActive {
				LoaderActive = false
				HelperTextDuration -= 1
				if HelperTextDuration <= 0 {
					projectIDLabel.Text = ""
					HelperTextActive = false
				}
			}

			if LoaderActive && !HelperTextActive {
				if ThemeMode == "Dark" {
					projectIDLabel.Color = color.RGBA{255, 255, 255, 255}
				} else {
					projectIDLabel.Color = color.RGBA{0, 0, 0, 255}
				}
				projectIDLabel.Alignment = fyne.TextAlignLeading

				if projectIDLabel.Text == "..." {
					projectIDLabel.Text = "... ..."
				} else if projectIDLabel.Text == "... ..." {
					projectIDLabel.Text = "... ... ..."
				} else if projectIDLabel.Text == "... ... ..." {
					projectIDLabel.Text = "..."
				} else {
					projectIDLabel.Text = "..."
				}
				projectIDLabel.Refresh()
			}
		}
	}
}

// ///////////////////////////
// Disable UI elements on start
func startDisUI() {
	endTimeOP.Hide()
	timeOptionLabel.Hide()
	startTimeOP.Hide()
	timeMinusOP.Hide()
	saveOptions.Hide()
	ApiPopButton.Hide()
	WebButton.Hide()

	timeStatText.Hide()

	apiKeyLabel.Hide()
	apiKeyEntry.Hide()

	urlLabel.Hide()
	urlEntry.Hide()
	projectIDEntry.Hide()

	sendToGitButton.Hide()
	sendToGitEntry.Hide()
	GitProgressBar.Hide()
	GitCheckMark.Hide()

	content.Hide()
	entry.Hide()
	image.Hide()
	PlayButton.Hide()
	RecordButton.Hide()
	selectEntry.Hide()
	refeshButton.Hide()
	summaryEntry.Hide()
	timeWorkText.Hide()
	timeLeftText.Hide()
	timeDoneText.Hide()
	PjLable.Hide()
	PjNavNeg.Hide()
	dividers[2].Hide()

	ShowTimeDayLabel.Hide()
	ShowTimeUserLabel.Hide()
	ShowTimeEntry.Hide()
	ShowTimeRefresh.Hide()
}

// //////////////////////////
// FILE HANDELING FUNCTIONS
func writeToLog(FormatedTime string) {
	var TEMP string
	logFile, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}
	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		TEMP += line + "\n"
	}
	gt := yct.Format("(2006-01-02) ")
	gt += FormatedTime
	TEMP += gt
	data := []byte(TEMP)
	gerr := os.WriteFile(Directory+"/log.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while saving Arrived to log")
	}
}
func newWriteToLog(Line string) {
	var TEMP string
	logFile, err := os.Open(Directory + "/log.txt")
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
	gerr := os.WriteFile(Directory+"/log.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while saving Arrived to log")
	}
}
func readFromLog(date time.Time, entry *widget.Entry) {
	entry.Text = ""
	tnf := date.Format("(2006-01-02)")
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

			if strings.Contains(new, "-|- Task:") {
				if strings.Contains(new, "-|-iid:") {
					new = " " + takeTextFromString(new, 0, strings.Index(new, "-|-iid:"))
				}
			}

			entry.MultiLine = true
			entry.Text += new + "\n"
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

// ///////////////////////////
// TASK RELATED FUNCTIONS
func GetSummaryFromTask(task string) string {
	//var sumHolder string
	fmt.Println("//////////////////////////////")
	fmt.Println(task)
	File, err := os.Open(Directory + "/response.txt")
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
				fmt.Println(line)
				return line
			}
		}
	}
	return "Error"
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
func MakeTasksLogable(selectEntryList []string, selectEntry *widget.SelectEntry) {
	selectEntryList = selectEntryList[:0]
	for _, issue := range issueHolder {
		selectEntryList = append(selectEntryList, issue.Title)
	}
	selectEntry.SetOptions(selectEntryList)
}
func MakeTasksLogable1(selectEntryList []string, selectEntry *widget.SelectEntry) {
	selectEntryList = selectEntryList[:0]
	for _, issue := range issueHolder {
		selectEntryList = append(selectEntryList, issue.Title)
	}
	selectEntry.SetOptions(selectEntryList)
}
func UpdateSelectProjectOptions() {
	for _, project := range ProjectIssuesHolder {
		ProjectList = append(ProjectList, project.ProjectName)
	}
	PjLable.SetOptions(ProjectList)
}
func FilterIssuesByProject() {
	var stringList []string
	var added bool
	var delTemp bool
	var currentTask string
	for _, issue := range issueHolder {
		added = false
		currentTask = issue.Title
		for round, project := range ProjectIssuesHolder {
			if issue.PJID == project.ProjectID {
				project.Tasks = append(project.Tasks, issue.Title)
				ProjectIssuesHolder[round].Tasks = append(ProjectIssuesHolder[round].Tasks, issue.Title)
				added = true
			}
		}
		if !added {
			if !delTemp {
				ProjectIssuesHolder = nil
				delTemp = true
			}
			stringList = append(stringList, currentTask)
			ProjectIssuesHolder = append(ProjectIssuesHolder, ProjectMap{issue.PJID, GetPJname(strconv.Itoa(issue.PJID)), stringList})
			stringList = nil
		}
	}
}
func GetNumOfTasks() {
	NumofTasks = 0
	File, err := os.Open(Directory + "/response.txt")
	if err != nil {
		panic(err)
	}

	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		line := Scanner.Text()
		clear0 := strings.Contains(line, "-|- Summary:")
		if clear0 {
			NumofTasks += 1
		}
	}
	TaskTakeValue = 1 / float64(NumofTasks)
}
func ValidateTask(list []string) bool {
	var NumberOfTasks int
	var NumberOfCorrectTasks int
	var Found bool

	for _, value := range list {
		if strings.Contains(value, "-|- Task:") {
			NumberOfTasks += 1
			VTask := takeTextFromString(value, strings.Index(value, "-|- Task:")+9, strings.Index(value, "-|- Summary:"))
			for _, value0 := range issueHolder {
				if VTask == value0.Title {
					NumberOfCorrectTasks += 1
					Found = true
				}
			}
			if !Found {
				for _, value1 := range LegacyTask {
					if VTask == value1 {
						NumberOfCorrectTasks += 1
						continue
					}
				}
			}
			Found = false
		}
	}
	if NumberOfTasks == NumberOfCorrectTasks {
		return true
	} else {
		return false
	}
}

var LegacyTask []string

func FetchLegacyTasks() {
	var foundTask bool
	LegacyTask = nil

	logFile, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		if strings.Contains(line, "-|- Task:") {
			VTask := takeTextFromString(line, strings.Index(line, "-|- Task:")+9, strings.Index(line, "-|- Summary:"))
			for _, val := range issueHolder {
				if VTask == val.Title {
					foundTask = true
				}
			}
			if !foundTask {
				LegacyTask = append(LegacyTask, VTask)
			}
			foundTask = false
		}

	}
}

// ///////////////////////////
// TEXT MODIFIKATION FUNCTIONS
func replaceSpaces(input string) string {
	replacedString := strings.ReplaceAll(input, " ", "+")
	return replacedString
}
func takeTextFromString(text string, start, end int) string {
	return strings.TrimSpace(text[start:end])
}

// ///////////////////////////
// SLICE FUNCTIONS
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

// ///////////////////////////
// TIME RELATED FUNCTIONS
func NavigateTooCurrentDay(yct time.Time, date *widget.Label, entry *widget.Entry, readyToSendLog []string, markButton *widget.Check, restrictMarkIcon *widget.Icon) {
	yct = time.Now()
	date.Text = yct.Weekday().String()
	date.Text += yct.Format(", 2006-,01-02 ")
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
	clearDate := fileContents(Directory+"/markedDates.txt", FormatedYCT10)
	if clearDate {
		restrictMarkIcon.Show()
		markButton.Disable()
	} else {
		restrictMarkIcon.Hide()
		markButton.Enable()
	}
}
func CheckEndTime(endTimeOP *widget.Entry) string {
	refDate := time.Now().Format("2006-01-02") // Use today's date as reference
	t, err := time.ParseInLocation("2006-01-02 15:04", refDate+" "+endTimeOP.Text, time.Local)
	if err != nil {
		return "Error"
	}
	endTime = t
	return t.Format("15:04")
}
func ChangeEndTime() {

	t, errd := time.Parse("15:04", "00:00")
	if errd != nil {
		fmt.Println("Error while reading from useroptions.txt")
	}
	DurationHold := workLength.Sub(t)
	endTime = arrivedTime.Add(DurationHold)
}
func checkArrivedStatus() bool {
	gg := time.Now()
	gt := gg.Format("(2006-01-02) ")

	File, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}
	var first bool
	var second bool
	scanner := bufio.NewScanner(File)
	for scanner.Scan() {
		line := scanner.Text()
		first = strings.Contains(line, gt)
		if first {
			second = strings.Contains(line, "-|- Arrived.")
			if second {
				return true
			}
		}
	}
	return false
}

// ///////////////////////////
// HELPER FUNCTIONS
func helpText(text string, Color string, length int) {
	projectIDLabel.Refresh()
	if Color == "blue" {
		projectIDLabel.Color = color.RGBA{30, 71, 232, 255}
	}
	if Color == "green" {
		projectIDLabel.Color = color.RGBA{0, 255, 0, 155}
	}
	if Color == "yellow" {
		projectIDLabel.Color = color.RGBA{230, 183, 55, 255}
	}
	if Color == "red" {
		projectIDLabel.Color = color.RGBA{255, 0, 0, 155}
	}
	projectIDLabel.Text = text
	projectIDLabel.Alignment = fyne.TextAlignCenter
	HelperTextActive = true
	HelperTextDuration = length * 2
}

var err error

func LoadInUserOptions() {
	// FETCH APPLICATIONS INFO FROM FILES
	var err error

	//var projectID string
	file, ferr := os.Open(Directory + "/useroptions.txt")
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
			workLength = endTime
		}
		// Load API KEY
		if i == 2 {
			ThemeMode = line
		}

		if i == 3 {
			URL = line
		}

		if i == 4 {
			Directory = line
		}
	}

	// clear response
	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile(Directory+"/response.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}
}

// ///////////////////////////
// SEND TO GIT FUNCTIONS
func loadReadyDatesToEntry(readyToSendLog []string, sendToGitEntry *widget.Entry) {
	sendToGitEntry.Text = ""
	sendToGitEntry.TextStyle.Bold = true

	TEMP9 := ""
	data := []byte(TEMP9)
	gerr := os.WriteFile(Directory+"/response.txt", data, 0644)
	if gerr != nil {
		fmt.Println("Error while writeing to Tasks")
	}

	sendToGitEntry.Text = "Press send button to log time to gitlab"
	for _, v := range readyToSendLog {
		v0 := v
		v += "\n"
		sendToGitEntry.Text += "\n"
		sendToGitEntry.Text += "-|- Date: " + v

		File, err := os.Open(Directory + "/log.txt")
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(File)

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

						summary0 := summary
						if strings.Contains(summary, "-|-iid:") {
							summary = takeTextFromString(summary, 0, strings.Index(summary, "-|-iid:"))
						}

						formatedTaskTime = strings.Replace(formatedTaskTime, summary0, "", -1)
					}
					// make delText == task name
					delText = strings.Replace(delText, summary, "", -1)
					delText = delText[2:]

					sendToGitEntry.MultiLine = true
					sendToGitEntry.Text += formatedTaskTime + "\n"
					sendToGitEntry.Text += summary + "\n" + "\n"
					writeToTasks("-|- Date: "+v+"\n", Directory+"/response.txt")
					writeToTasks(formatedTaskTime+"\n", Directory+"/response.txt")
					writeToTasks(delText+summary+"\n", Directory+"/response.txt")
					sendToGitEntry.Refresh()
				}
			}
		}

		sendToGitEntry.Refresh()
	}
}

// ///////////////////////////
// EXTRA
func restart() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Set up cmd command
	cmd := exec.Command(exe, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// starta nya iteration av programet
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	// Close det nu varande applicationen
	os.Exit(0)
}
