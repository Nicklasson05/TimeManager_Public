package main

import (
	"fmt"
	"image/color"
	"os"
	"os/user"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//var application = app.New()
var window = a.NewWindow("Time Manager Setup")

func setup() {
	//application.Settings().SetTheme(newTheme())
	window.Resize(fyne.NewSize(570, 210))

	window.SetContent(setupGui())
	window.ShowAndRun()
}

////////
//////// GUI section
////////
func setupGui() fyne.CanvasObject {
	// Varuabels

	// UI elements
	background0 := canvas.NewRectangle(color.RGBA{220, 220, 220, 255})
	background := canvas.NewRectangle(color.RGBA{200, 200, 200, 255})
	btn := widget.NewLabel("Enter where you would like Time Manager to be located?")
	dirEntry := widget.NewEntry()
	btn.TextStyle.Bold = true
	user, _ := user.Current()
	userDir := user.HomeDir
	dirEntry.Text = userDir

	saveBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		// Create directory on current user
		userDir = dirEntry.Text
		userDir += "/TimeManagerSaves"
		err := os.Mkdir(userDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
		// Create .txt files that is needed
		//issueIDHolder.txt
		file, err := os.Create(userDir + "/issueIDHolder.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file.Close()
		//issuePJIDHolder.txt
		file0, err := os.Create(userDir + "/issueIDHolder.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file0.Close()
		//log.txt
		file1, err := os.Create(userDir + "/log.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file1.Close()
		//markedDates.txt
		file2, err := os.Create(userDir + "/markedDates.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file2.Close()
		//response.txt
		file3, err := os.Create(userDir + "/response.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file3.Close()
		//responseIID.txt
		file4, err := os.Create(userDir + "/responseIID.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file4.Close()
		//responsePJID.txt
		file5, err := os.Create(userDir + "/responsePJID.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file5.Close()
		//Tasks.txt
		file6, err := os.Create(userDir + "/Tasks.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file6.Close()
		//useroptions.txt
		file7, err := os.Create(userDir + "/useroptions.txt")
		if err != nil {
			fmt.Println("Error creating .txt file:", err)
		}
		defer file7.Close()
		writeToTasks("16:00", userDir+"/useroptions.txt")
		writeToTasks("APIKEY", userDir+"/useroptions.txt")
		writeToTasks("TEMP", userDir+"/useroptions.txt")
		writeToTasks(userDir, userDir+"/useroptions.txt")
		//fmt.Println("Saved:" + userDir)

		// Carry the Dictionary path to start
		restart()
	})

	// Send to Layout
	newobj := []fyne.CanvasObject{background0, background, btn, dirEntry, saveBtn}
	return container.New(newkonstruktorLayout(background0, background, btn, dirEntry, saveBtn), newobj...)
}

////////
//////// Layout Section
////////
type newLayout struct {
	background0, background, btn, dirEntry, saveBtn fyne.CanvasObject
}

// MinSize implements fyne.Layout.
func (f *newLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	//border := fyne.NewSize(sideWidth*2, f.btn.MinSize().Height)
	border := fyne.NewSize(460, 110)
	return border.AddWidthHeight(100, 100)
}

func newkonstruktorLayout(background0, background, btn, dirEntry, saveBtn fyne.CanvasObject) fyne.Layout {
	return &newLayout{background0: background0, background: background, btn: btn, dirEntry: dirEntry, saveBtn: saveBtn}
}

func (f *newLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {

	f.saveBtn.Move(fyne.NewPos(470, 145))
	f.saveBtn.Resize(fyne.NewSize(50, 30))

	f.btn.Move(fyne.NewPos(65, 50))
	f.btn.Resize(fyne.NewSize(120, 30))

	f.dirEntry.Move(fyne.NewPos(67, 80))
	f.dirEntry.Resize(fyne.NewSize(420, 30))

	f.background.Move(fyne.NewPos(30, 30))
	f.background.Resize(fyne.NewSize(500, 150))

	f.background0.Move(fyne.NewPos(20, 20))
	f.background0.Resize(fyne.NewSize(520, 170))
}
