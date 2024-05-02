package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/canvas"
)

// summer time check var
var timeModifier int

// turn true to turn off LiveOverrideClock
var turnOffClock bool
var Magic bool

// CLOCK STATUS HOLDERS
var liveClockActive bool
var liveClockOverrideActive bool
var ClockActive bool

var endTimeUTD time.Time

var timeStatText *canvas.Text
var timeWorkText *canvas.Text
var timeLeftText *canvas.Text
var timeDoneText *canvas.Text

// status text varubels
var workTime time.Duration
var totalWorkTime time.Duration
var timeWork time.Duration
var shiftLeft time.Duration

func MakeStatusText() {
	// STATUS TEXTS

	// slacking
	timeStatText = canvas.NewText("Slacking: 0 h 0 min ", color.Black)
	timeStatText.TextSize = 15
	// WorkDone
	timeWorkText = canvas.NewText("Work Done: 0 h 0 min ", color.Black)
	timeWorkText.TextSize = 15
	// shift left
	timeLeftText = canvas.NewText("Shift Left: 0 h 0 min ", color.Black)
	timeLeftText.TextSize = 15
	// worked today
	timeDoneText = canvas.NewText("Worked Today: 0 h 0 min", color.Black)
	timeDoneText.TextSize = 15
}

func LiveClockOverride(ch chan int, quit chan bool, line string, shiftLeft time.Duration, endTime time.Time, timeLeftText *canvas.Text, arrivedTime time.Time, timeDoneText *canvas.Text) {
	tf := strings.Contains(line, "Arrived")
	ticker := time.NewTicker(1 * time.Second)
	var tikCounter int
	// fix summer time modifier
	if checkSummerTime() {
		timeModifier = -2
	} else {
		timeModifier = -1
	}

	// setting endtime up to date(UTD)
	currentTime := time.Now()
	clone := endTime
	clone10 := endTime
	clone10 = clone.Add(time.Duration(timeModifier) * time.Hour)
	endFormated := clone10.Format("15:04")
	currentFormated := currentTime.Format("2006-01-02 ")
	currentFormated += endFormated
	var err error
	endTimeUTD, err = time.Parse("2006-01-02 15:04", currentFormated)
	// Fixing arrived time
	clone0 := arrivedTime
	// make arrived time UTC
	cloneF0 := clone0.Format("15:04")
	CrtTimF := currentTime.Format("2006-01-02 ")
	CrtTimF += cloneF0
	clone0, _ = time.Parse("2006-01-02 15:04", CrtTimF)
	arrivedTime = clone0.Add(time.Duration(timeModifier) * time.Hour)

	if err != nil {
		fmt.Println("Error while trying to parse currentFormated in LiveClockOverride")
	}

	// RUNNING CLOCK
	if tf {
		for {
			select {
			case <-quit:
				liveClockOverrideActive = false
				return
			default:
				if Magic {
					t, errd := time.Parse("15:04", "00:00")
					if errd != nil {
						fmt.Println("Error while reading from useroptions.txt")
					}
					DurationHold := workLength.Sub(t)
					endTime = arrivedTime.Add(DurationHold)

					currentTime := time.Now()
					clone20 := endTime
					endFormated := clone20.Format("15:04")
					currentFormated := currentTime.Format("2006-01-02 ")
					currentFormated += endFormated
					endTimeUTD, _ = time.Parse("2006-01-02 15:04", currentFormated)
					Magic = false
				}

				// LOGIC
				currentTime = time.Now()
				ShiftLeft := endTimeUTD.Sub(currentTime)
				tikCounter += 1
				//fmt.Println(tikCounter)
				ClockActive = true
				// STOPPER
				if turnOffClock {
					turnOffClock = false
					ClockActive = false
					return
				}

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
				<-ticker.C
			}
		}
	} else {
		fmt.Println("NO Arrived exists in log")
	}
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
	return shiftLeft
}

func ResetStatusTexts(ch chan int, quit0 chan bool, quit chan bool, state bool) {
	//tn = time.Now()
	tnf = yct.Format("(2006-01-02)")

	shiftLeft = 0
	workTime = 0
	timeWork = 0

	liveClockOverrideActive = false
	liveClockActive = false

	logFile, err := os.Open(Directory + "/log.txt")
	if err != nil {
		panic(err)
	}

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		line := logScanner.Text()
		Clear := strings.Contains(line, tnf)
		if Clear {
			// fix arrived time
			arrivedTime = updateArrivedTime(line, arrivedTime)
			ChangeEndTime()
			// fix latest time
			latestArchivedTime = updateLatestArchivedTime(line, latestArchivedTime)

			// status text
			if state {
				shiftLeft = updateShiftLeftText(endTime)
			}
			workTime += updateWorkDoneStatus(line, totalWorkTime)

			arr := strings.Contains(line, "Arrived")
			if arr {
				arrActive = true
			}

			if !liveClockOverrideActive && state && !ClockActive {
				go LiveClockOverride(ch, quit0, line, shiftLeft, endTime, timeLeftText, arrivedTime, timeDoneText)
				liveClockOverrideActive = true
				turnOffClock = false
			} else if !state {
				turnOffClock = true
				liveClockActive = false
				liveClockOverrideActive = false
			} else if state && ClockActive && turnOffClock {
				//BUG fix
				turnOffClock = false
			}
		}

	}
	timeWork = workTime

	TEMP := "Work Done: "
	TEMP += fmt.Sprintf("%d h %d min", int(workTime.Hours()), int(workTime.Minutes())%60)
	timeWorkText.Text = TEMP
	timeWorkText.Refresh()

	if state {
		TEMP13 := "Shift Left: "
		TEMP13 += fmt.Sprintf("%d h %d min ", int(shiftLeft.Hours()), int(shiftLeft.Minutes())%60)
		timeLeftText.Text = TEMP13
		timeLeftText.Refresh()
	}

	if !state {
		timeLeftText.Text = "Shift Left: --- ---"
		timeDoneText.Text = "Worked Today: --- ---"
		timeLeftText.Refresh()
		timeDoneText.Refresh()
	}
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
func checkSummerTime() bool {
	year := time.Now().Year()
	date := time.Now()

	// hitta sista söndagen dagen i mars
	aprilFirst := time.Date(year, time.April, 1, 0, 0, 0, 0, time.UTC)
	daysToGoBack := int(aprilFirst.Weekday())
	lastSundayInMars := aprilFirst.AddDate(0, 0, -daysToGoBack)

	// hitta sista söndagen i october
	var LastSundayInOctober time.Time
	var Break bool
	for i := 31; i >= 1; i-- {
		currentDay := time.Date(year, time.October, i, 0, 0, 0, 0, time.UTC)
		if currentDay.Weekday() == time.Sunday && !Break {
			LastSundayInOctober = currentDay
			Break = true
		}
	}

	// kolla om det nu varande datumet kommer imellan de två datumen
	return date.After(lastSundayInMars) && date.Before(LastSundayInOctober)
}
