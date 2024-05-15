package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Options
var ApiPopButton *widget.Button
var endTimeOP *widget.Entry
var timeOptionLabel *widget.Label
var saveOptions *widget.Button
var urlEntry *widget.Entry
var urlLabel *canvas.Text
var apiKeyLabel *widget.Label
var apiKeyEntry *widget.Label
var WebButton *widget.Check

// options varubaels
var endTime time.Time
var workLength time.Time
var projectIDLabel *canvas.Text
var ThemeMode string
var ShowWebUrlBool bool

var startTimeOP *widget.Entry
var timeMinusOP *widget.Label
var projectIDEntry *widget.Label

func MakeOptions() {
	endTimeOP = widget.NewEntry()
	timeOptionLabel = widget.NewLabel("How long is your workday?")
	endTimeOP.Text = endTime.Format("15:04")

	urlEntry = widget.NewEntry()
	urlEntry.PlaceHolder = "URL"
	urlEntry.Text = URL
	urlLabel = canvas.NewText("Hello", color.RGBA{255, 0, 0, 255})

	apiKeyLabel = widget.NewLabel("Enter your gitlab url here")
	apiKeyEntry = widget.NewLabel("Re-enter Apikey")

	// WHAT TIME DO YOU END? OPTIONS
	startTimeOP = widget.NewEntry()
	timeMinusOP = widget.NewLabel("Hours")

	projectIDLabel = canvas.NewText("", color.RGBA{0, 255, 0, 155}) // NOW IS SAVED STATUS TEXT
	projectIDLabel.TextStyle.Bold = true
	projectIDLabel.Alignment = fyne.TextAlignCenter
	projectIDEntry = widget.NewLabel("")

	// Dark Mode
	PjNavNeg = widget.NewCheck("DarkMode", func(b bool) {
		setThemeMode(b)
	})
	// OPTION SAVE BUTTON
	saveOptions = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		// ENDTIME OPTION && API KEY
		LoaderActive = true
		var themeColor string
		timeValue := CheckEndTime(endTimeOP)
		if timeValue == "Error" {
			endTimeOP.Text = workLength.Format("15:04")
			endTimeOP.Refresh()
			helpText("workday incorrct", "red", 2)
		} else {
			if PjNavNeg.Checked {
				themeColor = "Dark"
			} else {
				themeColor = "Light"
			}

			if urlEntry.Text != "" {
				URL = urlEntry.Text
			}

			themeColor += "\n"

			workLength, err = time.Parse("15:04", timeValue)
			if err != nil {
				fmt.Println("Error while parsing timeValue")
			}
			Magic = true
			timeValue += "\n"
			data := timeValue + themeColor + URL
			newData := []byte(data)
			err := os.WriteFile(Directory+"/useroptions.txt", newData, 0644)
			if err != nil {
				fmt.Println("Error while saving options")
			}

			if checkSummerTime() {
				timeModifier = -2
			} else {
				timeModifier = -1
			}

			// setting endtime up to date(UTD)

			// reset tasks
			IDHolder = []IDStruct{}
			PJIDHolder = []IDStruct{}
			selectEntryList = []string{"Enter a valid API key in options"}
			selectEntry.SetOptions(selectEntryList)

			// get tasks
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

			// Reset status text
			turnOffClock = true
			FcurrentTime := time.Now().Format("(2006-01-02)")
			search := yct.Format("(2006-01-02)")
			if FcurrentTime != search {
				ResetStatusTexts(ch, quit0, quit, false)
			} else {
				ResetStatusTexts(ch, quit0, quit, true)
			}
			helpText("SAVED", "green", 1)
		}
	})

	ApiPopButton = widget.NewButton("Apikey", func() {
		EnablePop()
	})

	WebButton = widget.NewCheck("Show Task Url", func(b bool) {
		if b {
			ShowWebUrlBool = true
			ShowWebUrl()
		} else {
			ShowWebUrlBool = false
			for index, issue := range ProjectIssuesHolder {
				if PjLable.Selected == issue.ProjectName {
					if ShowWebUrlBool {
						ShowWebUrl()
					} else {
						selectEntry.SetOptions(ProjectIssuesHolder[index].Tasks)
					}
				}
			}
		}
	})
}
