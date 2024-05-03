package main

import (
	"os/user"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// APPLICATION DECLARATION
var a = app.New()
var w = a.NewWindow("Time Manager")

var Directory string

type IDStruct struct {
	ID, Task string
}

var PJIDHolder []IDStruct
var IDHolder []IDStruct
var UserID string

func main() {

	// GENERAL SETUP
	a.Settings().SetTheme(newTheme())
	w.Resize(fyne.NewSize(700, 400))

	//Icon
	a.SetIcon(resourceIcon)
	w.SetIcon(resourceIcon)

	// SETUP
	user, _ := user.Current()
	userDir := user.HomeDir
	setupDone := SetupDone(userDir, "TimeManagerSaves")

	if setupDone {
		Directory = userDir + "/TimeManagerSaves"
		w.SetContent(makeGUI())
		w.ShowAndRun()
	} else {
		setup()
	}
}

// GENERAL FUNCTONS
func NewSize(i1, i2 int) {
	panic("unimplemented")
}
func UppdateUI() {
	w.Canvas().Refresh(makeGUI())
}
