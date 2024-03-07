package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 220

type nyLayout struct {
	top, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, summaryEntry, content, image fyne.CanvasObject
	dividers                                                                                                                                                                                                                                                                                                                              [3]fyne.CanvasObject
}

func nykonstruktorLayout(top, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, summaryEntry, content, image fyne.CanvasObject, dividers [3]fyne.CanvasObject) fyne.Layout {
	return &nyLayout{top: top, date: date, sendToGitEntry: sendToGitEntry, sendToGitButton: sendToGitButton, markButton: markButton, restrictMarkIcon: restrictMarkIcon, projectIDEntry: projectIDEntry, projectIDLabel: projectIDLabel, apiKeyEntry: apiKeyEntry, apiKeyLabel: apiKeyLabel, toolbarSide: toolbarSide, timeStatText: timeStatText, timeWorkText: timeWorkText, timeLeftText: timeLeftText, timeDoneText: timeDoneText, timeOptionLabel: timeOptionLabel, startTimeOP: startTimeOP, endTimeOP: endTimeOP, timeMinusOP: timeMinusOP, saveOptions: saveOptions, sidebar: sidebar, entry: entry, selectEntry: selectEntry, summaryEntry: summaryEntry, content: content, dividers: dividers, image: image}
}

func (l *nyLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	// DATE NAVIGATION

	l.markButton.Move(fyne.NewPos(280, 3))
	l.markButton.Resize(fyne.NewSize(30, 30))

	l.restrictMarkIcon.Move(fyne.NewPos(280, 3))
	l.restrictMarkIcon.Resize(fyne.NewSize(30, 30))

	l.toolbarSide.Move(fyne.NewPos(size.Width-120, 0))
	l.toolbarSide.Resize(fyne.NewSize(size.Width, topHeight))

	l.date.Move(fyne.NewPos(120, 2))
	l.date.Resize(fyne.NewSize(sideWidth, size.Height-topHeight))

	l.image.Move(fyne.NewPos(size.Width-250, topHeight+5))
	l.image.Resize(fyne.NewSize(25, 25))

	// SEND MENU

	l.sendToGitButton.Move(fyne.NewPos(size.Width-sideWidth+168, topHeight))
	l.sendToGitButton.Resize(fyne.NewSize(50, 25))

	l.sendToGitEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+50))
	l.sendToGitEntry.Resize(fyne.NewSize(205, 155))

	// LOG MENU

	l.selectEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+50))
	l.selectEntry.Resize(fyne.NewSize(205, 35))

	l.summaryEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+100))
	l.summaryEntry.Resize(fyne.NewSize(205, 35))

	l.entry.Move(fyne.NewPos(-5, topHeight))
	l.entry.Resize(fyne.NewSize(size.Width-sideWidth*2+225, size.Height-topHeight))

	l.sidebar.Move(fyne.NewPos(size.Width-sideWidth+35, topHeight))
	l.sidebar.Resize(fyne.NewSize(80, 80-topHeight))

	// OPTIONS MENU

	l.timeOptionLabel.Move(fyne.NewPos(size.Width-sideWidth+35, topHeight+20))

	//+ button
	l.startTimeOP.Move(fyne.NewPos(size.Width-sideWidth+140, topHeight+50))
	l.startTimeOP.Resize(fyne.NewSize(50, 35))
	//- button
	l.timeMinusOP.Move(fyne.NewPos(size.Width-sideWidth+30, topHeight+50))
	l.timeMinusOP.Resize(fyne.NewSize(50, 35))

	l.endTimeOP.Move(fyne.NewPos(size.Width-sideWidth+80, topHeight+50))
	l.endTimeOP.Resize(fyne.NewSize(60, 35))

	// API KEY OPTION
	l.apiKeyLabel.Move(fyne.NewPos(size.Width-sideWidth+18, topHeight+80))
	l.apiKeyEntry.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+110))
	l.apiKeyEntry.Resize(fyne.NewSize(170, 35))

	// Project ID OPTION
	l.projectIDLabel.Move(fyne.NewPos(size.Width-sideWidth+50, topHeight-26))

	l.projectIDEntry.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+170))
	l.projectIDEntry.Resize(fyne.NewSize(170, 35))

	//Save Button
	l.saveOptions.Move(fyne.NewPos(size.Width-sideWidth+168, topHeight))
	l.saveOptions.Resize(fyne.NewSize(50, 25))

	//Status texts
	l.timeDoneText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-80))
	l.timeDoneText.Resize(fyne.NewSize(100, 100-topHeight))

	l.timeWorkText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-140))
	l.timeWorkText.Resize(fyne.NewSize(100, 100-topHeight))

	l.timeStatText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-120))
	l.timeStatText.Resize(fyne.NewSize(100, 100-topHeight))

	l.timeLeftText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-100))
	l.timeLeftText.Resize(fyne.NewSize(100, 100-topHeight))

	// DESIGNER
	l.content.Move(fyne.NewPos(sideWidth-300, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sideWidth*2+300, size.Height-topHeight))

	dividerThickness := theme.SeparatorThicknessSize()
	l.dividers[0].Move(fyne.NewPos(0, topHeight))
	l.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	l.dividers[2].Move(fyne.NewPos(size.Width-sideWidth, topHeight))
	l.dividers[2].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))

}

func (l *nyLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	//border := fyne.NewSize(sideWidth*2, l.top.MinSize().Height)
	border := fyne.NewSize(440, 180)
	return border.AddWidthHeight(100, 100)
}
