package main

import (
	"fmt"
	"image/color"
	"net/url"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// pop
var popBackBackground *canvas.Rectangle
var popBackground *canvas.Rectangle

var popEntry *widget.Entry
var popLink *widget.Hyperlink
var popLable *canvas.Text
var popIcon *widget.Icon
var popButton *widget.Button
var popUrlEntry *widget.Entry

var popActive bool

func MakePop() {
	// Enter API background

	popBackground = canvas.NewRectangle(color.Gray{Y: 0xee})
	popBackBackground = canvas.NewRectangle(color.RGBA{204, 204, 204, 255})
	popEntry = widget.NewPasswordEntry()
	popEntry.PlaceHolder = "Enter apikey here"
	tempUrl := URL

	var urlLink *url.URL
	if strings.Contains(tempUrl, "/api/") {
		index := strings.Index(tempUrl, "/api/")
		var err error
		urlLink, err = url.Parse(takeTextFromString(tempUrl, 0, index) + "/-/user_settings/personal_access_tokens")
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("URl is not API URL")
	}
	popLink = widget.NewHyperlink("Link to gitlab", urlLink)
	popLable = canvas.NewText("Get Access To Tasks", color.RGBA{100, 100, 100, 255})
	popLable.TextStyle.Bold = true
	popLable.TextSize = 13

	popIcon = widget.NewIcon(theme.WarningIcon())

	popButton = widget.NewButton("OK", func() {
		LoaderActive = true
		if !CheckAPIKEY(popEntry.Text) {
			helpText("Invalid APIKEY", "red", 2)
		} else {
			apiKey = popEntry.Text
			// API RELATED VARUABELS
			userID := RequestUserID(apiKey)
			RequestIssues(apiKey, userID)
			selectEntry.SetOptions(ProjectIssuesHolder[0].Tasks)
			PjLable.Selected = ProjectIssuesHolder[0].ProjectName
			DisablePop()
			helpText("ApiKey Accepted", "green", 1)
		}
	})
	popActive = true
}

func DisablePop() {
	popBackground.Hide()
	popBackBackground.Hide()
	popEntry.Hide()
	popLink.Hide()
	popLable.Hide()
	popIcon.Hide()
	popButton.Hide()
	popActive = false

	endTimeOP.Hide()
	timeOptionLabel.Hide()
	saveOptions.Hide()
	ApiPopButton.Hide()
	WebButton.Hide()
	PjNavNeg.Hide()

	apiKeyLabel.Hide()
	apiKeyEntry.Hide()

	urlEntry.Hide()

	content.Show()
	entry.Show()
	image.Show()
	PlayButton.Show()
	RecordButton.Show()
	selectEntry.Show()
	refeshButton.Show()
	summaryEntry.Show()
	timeWorkText.Show()
	timeLeftText.Show()
	timeDoneText.Show()
	PjLable.Show()
	dividers[2].Show()
}

func EnablePop() {
	popBackground.Show()
	popBackBackground.Show()
	popEntry.Show()
	popLink.Show()
	popLable.Show()
	popIcon.Show()
	popButton.Show()
	popEntry.Text = ""
	popEntry.Refresh()
	popActive = true

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

	endTimeOP.Hide()
	timeOptionLabel.Hide()
	saveOptions.Hide()
	urlEntry.Hide()
	urlLabel.Hide()
	ApiPopButton.Hide()
	apiKeyLabel.Hide()
	apiKeyEntry.Hide()
	WebButton.Hide()
}

// ////////////////////////////
// Web Url FUNCTIONS
func ShowWebUrl() {
	var Combo []string
	for index, _ := range ProjectIssuesHolder {
		if ProjectIssuesHolder[index].ProjectName == PjLable.Selected {
			for _, v0 := range ProjectIssuesHolder[index].Tasks {
				for _, v1 := range issueHolder {
					if v0 == v1.Title {
						Combo = append(Combo, v0+"    "+v1.WebUrl)
					}
				}
			}
		}
	}
	selectEntry.SetOptions(Combo)
}
