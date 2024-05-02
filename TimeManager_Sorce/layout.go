package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 220

var popLocationX int
var popLocationY int

var changeSize int

type nyLayout struct {
	top, popBackBackground, popBackground, popIcon, popEntry, popButton, popLink, popLable, ShowTimeEntry, ShowTimeDayLabel, ShowTimeUserLabel, ShowTimeRefresh, PlayButton, RecordButton, PjLable, PjNavNeg, GitProgressBar, GitCheckMark, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, urlEntry, urlLabel, ApiPopButton, WebButton, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, refeshButton, summaryEntry, content, image fyne.CanvasObject
	dividers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            [3]fyne.CanvasObject
}

func nykonstruktorLayout(popBackBackground, popBackground, popIcon, popEntry, popButton, popLink, popLable, top, ShowTimeEntry, ShowTimeDayLabel, ShowTimeUserLabel, ShowTimeRefresh, PlayButton, RecordButton, PjLable, PjNavNeg, GitProgressBar, GitCheckMark, date, sendToGitEntry, sendToGitButton, markButton, restrictMarkIcon, projectIDEntry, urlEntry, urlLabel, ApiPopButton, WebButton, projectIDLabel, apiKeyEntry, apiKeyLabel, toolbarSide, timeStatText, timeWorkText, timeLeftText, timeDoneText, timeOptionLabel, startTimeOP, endTimeOP, timeMinusOP, saveOptions, sidebar, entry, selectEntry, refeshButton, summaryEntry, content, image fyne.CanvasObject, dividers [3]fyne.CanvasObject) fyne.Layout {
	return &nyLayout{top: top, popBackBackground: popBackBackground, popBackground: popBackground, popIcon: popIcon, popEntry: popEntry, popButton: popButton, popLink: popLink, popLable: popLable, ShowTimeEntry: ShowTimeEntry, ShowTimeDayLabel: ShowTimeDayLabel, ShowTimeUserLabel: ShowTimeUserLabel, ShowTimeRefresh: ShowTimeRefresh, PlayButton: PlayButton, RecordButton: RecordButton, PjLable: PjLable, PjNavNeg: PjNavNeg, GitProgressBar: GitProgressBar, GitCheckMark: GitCheckMark, date: date, sendToGitEntry: sendToGitEntry, sendToGitButton: sendToGitButton, markButton: markButton, restrictMarkIcon: restrictMarkIcon, projectIDEntry: projectIDEntry, urlEntry: urlEntry, urlLabel: urlEntry, ApiPopButton: ApiPopButton, WebButton: WebButton, projectIDLabel: projectIDLabel, apiKeyEntry: apiKeyEntry, apiKeyLabel: apiKeyLabel, toolbarSide: toolbarSide, timeStatText: timeStatText, timeWorkText: timeWorkText, timeLeftText: timeLeftText, timeDoneText: timeDoneText, timeOptionLabel: timeOptionLabel, startTimeOP: startTimeOP, endTimeOP: endTimeOP, timeMinusOP: timeMinusOP, saveOptions: saveOptions, sidebar: sidebar, entry: entry, selectEntry: selectEntry, refeshButton: refeshButton, summaryEntry: summaryEntry, content: content, dividers: dividers, image: image}
}

func (l *nyLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight-2))

	// Pop
	popLocationX = int(size.Width)/2 - 110
	popLocationY = -70

	l.popBackground.Move(fyne.NewPos(12+float32(popLocationX), topHeight+150+float32(popLocationY)))
	l.popBackground.Resize(fyne.NewSize(200, 140-topHeight))

	l.popBackBackground.Move(fyne.NewPos(9+float32(popLocationX), topHeight+147+float32(popLocationY)))
	l.popBackBackground.Resize(fyne.NewSize(206, 146-topHeight))

	l.popEntry.Move(fyne.NewPos(20+float32(popLocationX), topHeight+175+float32(popLocationY)))
	l.popEntry.Resize(fyne.NewSize(180, 35))

	l.popButton.Move(fyne.NewPos(138+float32(popLocationX), topHeight+228+float32(popLocationY)))
	l.popButton.Resize(fyne.NewSize(70, 20))

	l.popLink.Move(fyne.NewPos(20+float32(popLocationX), topHeight+205+float32(popLocationY)))
	l.popLink.Resize(fyne.NewSize(90, 32))

	l.popLable.Move(fyne.NewPos(50+float32(popLocationX), topHeight+145+float32(popLocationY)))
	l.popLable.Resize(fyne.NewSize(100, 32))

	l.popIcon.Move(fyne.NewPos(20+float32(popLocationX), topHeight+152+float32(popLocationY)))
	l.popIcon.Resize(fyne.NewSize(20, 20))

	// DATE NAVIGATION
	l.markButton.Move(fyne.NewPos(280, 3))
	l.markButton.Resize(fyne.NewSize(30, 30))

	l.restrictMarkIcon.Move(fyne.NewPos(280, 3))
	l.restrictMarkIcon.Resize(fyne.NewSize(30, 30))

	//l.toolbarSide.Move(fyne.NewPos(size.Width-120, -2))
	//l.toolbarSide.Resize(fyne.NewSize(size.Width, topHeight))

	l.toolbarSide.Move(fyne.NewPos(size.Width-160, -2))
	l.toolbarSide.Resize(fyne.NewSize(size.Width, topHeight))

	l.date.Move(fyne.NewPos(120, 2))
	l.date.Resize(fyne.NewSize(sideWidth, size.Height-topHeight))

	l.image.Move(fyne.NewPos(size.Width-285, topHeight+5))
	l.image.Resize(fyne.NewSize(60, 25))

	// SEND MENU
	l.sendToGitButton.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+35))
	l.sendToGitButton.Resize(fyne.NewSize(80, 25))

	l.GitCheckMark.Move(fyne.NewPos(size.Width-sideWidth+110, topHeight+35))
	l.GitCheckMark.Resize(fyne.NewSize(60, 25))

	//l.sendToGitEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+30))
	//l.sendToGitEntry.Resize(fyne.NewSize(205, 175))

	l.sendToGitEntry.Move(fyne.NewPos(-5, topHeight))
	l.sendToGitEntry.Resize(fyne.NewSize(size.Width-sideWidth*2+225, size.Height-topHeight))

	l.GitProgressBar.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+5))
	l.GitProgressBar.Resize(fyne.NewSize(205, 25))

	// LOG MENU

	l.selectEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+95))
	l.selectEntry.Resize(fyne.NewSize(180, 35))

	//l.summaryEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+95))
	//l.summaryEntry.Resize(fyne.NewSize(205, 35))

	l.summaryEntry.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+135))
	l.summaryEntry.Resize(fyne.NewSize(205, 35))

	l.refeshButton.Move(fyne.NewPos(size.Width-sideWidth+193, topHeight+95))
	l.refeshButton.Resize(fyne.NewSize(23, 35))

	l.entry.Move(fyne.NewPos(-5, topHeight))
	l.entry.Resize(fyne.NewSize(size.Width-sideWidth*2+225, size.Height-topHeight))

	// Project Nav
	l.PjLable.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+55))
	l.PjLable.Resize(fyne.NewSize(205, 35))

	l.PjNavNeg.Move(fyne.NewPos(size.Width-sideWidth+55, topHeight+210))
	l.PjNavNeg.Resize(fyne.NewSize(23, 25))

	// Log buttons

	l.sidebar.Move(fyne.NewPos(size.Width-sideWidth+50, topHeight+150))
	l.sidebar.Resize(fyne.NewSize(80, 80-topHeight))

	l.PlayButton.Move(fyne.NewPos(size.Width-sideWidth+10, topHeight+5))
	l.PlayButton.Resize(fyne.NewSize(100, 80-topHeight))

	l.RecordButton.Move(fyne.NewPos(size.Width-sideWidth+115, topHeight+5))
	l.RecordButton.Resize(fyne.NewSize(100, 80-topHeight))

	// OPTIONS MENU
	l.timeOptionLabel.Move(fyne.NewPos(size.Width-sideWidth+35, topHeight+20))

	// WEB OPTION
	l.WebButton.Move(fyne.NewPos(size.Width-sideWidth+55, topHeight+235))
	l.WebButton.Resize(fyne.NewSize(100, 25))

	// Minuts entry
	l.startTimeOP.Move(fyne.NewPos(size.Width-sideWidth+140, topHeight+50))
	l.startTimeOP.Resize(fyne.NewSize(50, 35))
	//Minuts lable

	// Hours lable
	l.timeMinusOP.Move(fyne.NewPos(size.Width-sideWidth+30, topHeight+50))
	l.timeMinusOP.Resize(fyne.NewSize(50, 35))
	// Hours entry
	l.endTimeOP.Move(fyne.NewPos(size.Width-sideWidth+85, topHeight+50))
	l.endTimeOP.Resize(fyne.NewSize(49, 35))

	//Refresh apikey
	l.ApiPopButton.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+165))
	l.ApiPopButton.Resize(fyne.NewSize(170+float32(changeSize), 35))

	// url Lable
	l.apiKeyLabel.Move(fyne.NewPos(size.Width-sideWidth+32, topHeight+80))
	// api Lable
	l.apiKeyEntry.Move(fyne.NewPos(size.Width-sideWidth+58, topHeight+135))

	// Project ID OPTION
	//l.projectIDLabel.Move(fyne.NewPos(size.Width-sideWidth-40, topHeight-28))
	l.projectIDLabel.Move(fyne.NewPos(size.Width-sideWidth-30, topHeight-28))

	l.projectIDEntry.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+170))
	l.projectIDEntry.Resize(fyne.NewSize(170, 35))

	// URL OPTION
	//l.urlEntry.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+170))
	//l.urlEntry.Resize(fyne.NewSize(170, 35))
	l.urlEntry.Move(fyne.NewPos(size.Width-sideWidth+25, topHeight+105))
	l.urlEntry.Resize(fyne.NewSize(170, 35))

	//Save Button
	l.saveOptions.Move(fyne.NewPos(size.Width-sideWidth+158, topHeight))
	l.saveOptions.Resize(fyne.NewSize(60, 25))

	// ShowTime Menu

	l.ShowTimeEntry.Move(fyne.NewPos(-5, topHeight))
	l.ShowTimeEntry.Resize(fyne.NewSize(size.Width-sideWidth*2+225, size.Height-topHeight))

	l.ShowTimeRefresh.Move(fyne.NewPos(size.Width-295, topHeight+5))
	l.ShowTimeRefresh.Resize(fyne.NewSize(70, 25))

	l.ShowTimeDayLabel.Move(fyne.NewPos(size.Width-sideWidth+115, topHeight+5))
	//l.ShowTimeDayLabel.Resize(fyne.NewSize(100, 80-topHeight))

	l.ShowTimeUserLabel.Move(fyne.NewPos(size.Width-sideWidth+115, topHeight+50))
	//l.ShowTimeDayLabel.Resize(fyne.NewSize(100, 80-topHeight))

	//Status texts
	l.timeDoneText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-80))
	l.timeDoneText.Resize(fyne.NewSize(100, 100-topHeight))

	l.timeWorkText.Move(fyne.NewPos(size.Width-sideWidth+10, +w.Canvas().Size().Height-120))
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
	border := fyne.NewSize(440, 200)
	return border.AddWidthHeight(100, 100)
}
