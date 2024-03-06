package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// APPLICATION DECLARATION
var a = app.New()
var w = a.NewWindow("Time Manager")

func main() {
	// GENERAL SETUP
	a.Settings().SetTheme(newTheme())
	w.Resize(fyne.NewSize(700, 400))

	// SET ICON
	r, err := LoadResourceFromPath("TimeManagerIcon.png")
	if err != nil {
		fmt.Println(err)
	}
	w.SetIcon(r)

	// SHOW UI ELEMENTS + LOGIC
	w.SetContent(makeGUI())
	w.ShowAndRun()

}

// GENERAL FUNCTONS

func NewSize(i1, i2 int) {
	panic("unimplemented")
}
func UppdateUI() {
	w.Canvas().Refresh(makeGUI())
}

// ICON CODE

func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string {
	return r.StaticName
}

type Resource interface {
	Name() string
	Content() []byte
}

func (r *StaticResource) Content() []byte {
	return r.StaticContent
}
func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}
