//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type nyTheme struct {
	fyne.Theme
}

func newTheme() fyne.Theme {
	return &nyTheme{Theme: theme.DefaultTheme()}
}

func (t *nyTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func (t *nyTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}
