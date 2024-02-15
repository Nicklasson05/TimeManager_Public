package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var a = app.New()
var w = a.NewWindow("Time Manager")

func main() {
	//a := app.New()
	a.Settings().SetTheme(newTheme())
	//w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(700, 400))
	//w.SetFixedSize(true)
	//Content
	/*
		w.SetContent(widget.NewLabel("Det Fungerade!!!"))

		testButton := widget.NewButton("click me", func() {
			w.SetContent(widget.NewLabel("Button Clicked"))
		})
		testButton.Resize(fyne.NewSize(50, 25))

		content := container.New(layout.NewVBoxLayout(), testButton)
		w.SetContent(content)

	*/
	w.SetContent(makeGUI())
	w.ShowAndRun()

}

func NewSize(i1, i2 int) {
	panic("unimplemented")
}

func UppdateUI() {
	w.Canvas().Refresh(makeGUI())
}
