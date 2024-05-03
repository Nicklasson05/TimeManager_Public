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
func newTheme1() fyne.Theme {
	return &nyTheme{Theme: theme.DarkTheme()}
}

func setThemeMode(b bool) {
	if b {
		a.Settings().SetTheme(newTheme1())
		content.FillColor = color.RGBA{0, 0, 0, 0}
		timeStatText.Color = color.RGBA{255, 255, 255, 255}
		timeWorkText.Color = color.RGBA{255, 255, 255, 255}
		timeLeftText.Color = color.RGBA{255, 255, 255, 255}
		timeDoneText.Color = color.RGBA{255, 255, 255, 255}
		popLable.Color = color.RGBA{255, 255, 255, 255}
		popBackground.FillColor = color.RGBA{50, 50, 50, 255}
		ShowTimeDayLabel.Color = color.RGBA{255, 255, 255, 255}
		ShowTimeUserLabel.Color = color.RGBA{255, 255, 255, 255}
		ThemeMode = "Dark"
	} else {
		a.Settings().SetTheme(newTheme())
		content.FillColor = color.RGBA{0, 0, 0, 0}
		timeStatText.Color = color.RGBA{0, 0, 0, 255}
		timeWorkText.Color = color.RGBA{0, 0, 0, 255}
		timeLeftText.Color = color.RGBA{0, 0, 0, 255}
		timeDoneText.Color = color.RGBA{0, 0, 0, 255}
		popLable.Color = color.RGBA{100, 100, 100, 255}
		popBackground.FillColor = color.Gray{Y: 0xee}
		ThemeMode = "Light"
	}
}

// theme.DefaultTheme()
func (t *nyTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func (t *nyTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}
