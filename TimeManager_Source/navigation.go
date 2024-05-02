package main

import (
	"time"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var systemCheckChan chan int
var qsystemCheckChanExit chan bool
var quit chan bool
var ch chan int
var quit0 chan bool

// Time varuables
var tnf string
var yct time.Time
var ct time.Time

var toolbar *widget.Toolbar
var restrictMarkIcon *widget.Icon
var markButton *widget.Check
var date *widget.Label

func MakeNavigation() {
	date = widget.NewLabel("")
	ct = time.Now()
	date.Text = ct.Weekday().String()
	date.Text += ct.Format(", 2006-01-02 ")

	//System check channel
	systemCheckChan = make(chan int)
	qsystemCheckChanExit = make(chan bool)

	//c := make(chan int)
	quit = make(chan bool)

	ch = make(chan int)
	quit0 = make(chan bool)
	// CLOCK GOROUTINE
	yct = ct

	// MARK BUTTON CODE
	markButton = widget.NewCheck("", func(b bool) {
		FormatedYCT10 := yct.Format("(2006-01-02)")
		clearDate := fileContents(Directory+"/markedDates.txt", FormatedYCT10)
		if !clearDate {
			if b {
				FormatedYCT := yct.Format("(2006-01-02)")
				readyToSendLog = append(readyToSendLog, FormatedYCT)
				loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
				GetNumOfTasks()
			} else if !b {
				FormatedYCT0 := yct.Format("(2006-01-02)")
				readyToSendLog = removeElementFromList(readyToSendLog, FormatedYCT0)
				readyToSendLog = removeDuplicates(readyToSendLog)
				loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
				GetNumOfTasks()
			}
		}
	})

	// MARKED DATES CHECK
	FormatedYCT10 := yct.Format("(2006-01-02)")
	clearDate := fileContents(Directory+"/markedDates.txt", FormatedYCT10)
	if clearDate {
		restrictMarkIcon.Show()
		markButton.Disable()
	} else {
		restrictMarkIcon.Hide()
		markButton.Enable()
	}

	// DATE NAVIGATION
	toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			yct = yct.Add(-24 * time.Hour)
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()
			readFromLog(yct, entry)

			turnOffClock = true
			FcurrentTime := time.Now().Format("(2006-01-02)")
			search := yct.Format("(2006-01-02)")
			if FcurrentTime != search {
				ResetStatusTexts(ch, quit0, quit, false)
			} else {
				ResetStatusTexts(ch, quit0, quit, true)
			}

			// Check Box
			var exists bool
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
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			yct = yct.Add(+24 * time.Hour)
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()
			readFromLog(yct, entry)

			FcurrentTime := time.Now().Format("(2006-01-02)")
			search := yct.Format("(2006-01-02)")
			turnOffClock = true
			if FcurrentTime != search {
				ResetStatusTexts(ch, quit0, quit, false)
			} else {
				ResetStatusTexts(ch, quit0, quit, true)
			}

			// Check Box
			var exists bool
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

		}),
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			yct = time.Now()
			date.Text = yct.Weekday().String()
			date.Text += yct.Format(", 2006-01-02 ")
			date.Refresh()
			readFromLog(yct, entry)

			turnOffClock = true
			FcurrentTime := time.Now().Format("(2006-01-02)")
			search := yct.Format("(2006-01-02)")
			if FcurrentTime != search {
				ResetStatusTexts(ch, quit0, quit, false)
			} else {
				ResetStatusTexts(ch, quit0, quit, true)
			}

			// Check Box
			var exists bool
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

		}),
	)
}
