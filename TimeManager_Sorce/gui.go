package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// UI ELEMENTS
var dividers = [3]fyne.CanvasObject{
	widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(),
}

var content *canvas.Rectangle

func makeGUI() fyne.CanvasObject {

	// UI ELEMENTS
	content = canvas.NewRectangle(color.Gray{Y: 0xee})

	MakeSendToGit() // Git

	MakeStatusText() // status

	MakeNavigation() // Nav

	MakeLog() // Log

	go SystemCheck(systemCheckChan, qsystemCheckChanExit, PlayButton, RecordButton) //SystemCheck

	LoadInUserOptions() // Load In Options

	MakeOptions() // Options

	MakeMenu() // Menu

	MakeShowTime() // ShowTime

	LoadInLog() // Load In Log

	MakePop() // Pop

	startDisUI() // Hide UI on start

	if ThemeMode == "Dark" {
		setThemeMode(true)
		PjNavNeg.Checked = true
	} else {
		setThemeMode(false)
	}

	// SEND UI ELEMENTS TO GET POSITIONAL AND SIZE DATA TO layout.go
	obj := []fyne.CanvasObject{content, popBackBackground, popBackground, popIcon, popEntry, popButton, popLink, popLable, ShowTimeEntry, ShowTimeDayLabel, ShowTimeUserLabel, ShowTimeRefresh, PlayButton, RecordButton, PjLable, PjNavNeg, GitProgressBar, GitCheckMark, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, urlEntry, urlLabel, ApiPopButton, WebButton, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, toolbar, sidebar, entry, selectEntry, refeshButton, summaryEntry, dividers[0], dividers[1], dividers[2], image}
	return container.New(nykonstruktorLayout(popBackBackground, popBackground, popIcon, popEntry, popButton, popLink, popLable, toolbar, ShowTimeEntry, ShowTimeDayLabel, ShowTimeUserLabel, ShowTimeRefresh, PlayButton, RecordButton, PjLable, PjNavNeg, GitProgressBar, GitCheckMark, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, urlEntry, urlLabel, ApiPopButton, WebButton, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, refeshButton, summaryEntry, content, image, dividers), obj...)
}
