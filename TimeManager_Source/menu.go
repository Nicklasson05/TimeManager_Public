package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var toolbarSide *widget.Toolbar
var menu string

func MakeMenu() {
	// MENU SELECT
	toolbarSide = widget.NewToolbar(
		// LOG MENU
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			if !popActive {
				menu = "Log"
				sidebar.Show()
				selectEntry.Show()
				summaryEntry.Show()
				refeshButton.Show()
				PlayButton.Show()
				RecordButton.Show()
				entry.Show()
				image.Show()

				PjLable.Show()

				endTimeOP.Hide()
				timeOptionLabel.Hide()
				startTimeOP.Hide()
				timeMinusOP.Hide()
				saveOptions.Hide()
				ApiPopButton.Hide()
				PjNavNeg.Hide()
				WebButton.Hide()

				apiKeyLabel.Hide()
				apiKeyEntry.Hide()

				urlEntry.Hide()

				projectIDEntry.Hide()

				sendToGitButton.Hide()
				sendToGitEntry.Hide()
				GitCheckMark.Hide()
				GitProgressBar.Hide()

				ShowTimeDayLabel.Hide()
				ShowTimeUserLabel.Hide()
				ShowTimeEntry.Hide()
				ShowTimeRefresh.Hide()
			}
		}),
		// SEND MENU
		widget.NewToolbarAction(theme.MailComposeIcon(), func() {
			if !popActive {
				menu = "Git"
				sidebar.Hide()
				selectEntry.Hide()
				summaryEntry.Hide()
				refeshButton.Hide()
				PlayButton.Hide()
				RecordButton.Hide()
				entry.Hide()
				image.Hide()

				PjLable.Hide()
				PjNavNeg.Hide()

				endTimeOP.Hide()
				timeOptionLabel.Hide()
				startTimeOP.Hide()
				timeMinusOP.Hide()
				saveOptions.Hide()
				ApiPopButton.Hide()
				WebButton.Hide()

				apiKeyLabel.Hide()
				apiKeyEntry.Hide()

				urlEntry.Hide()

				projectIDEntry.Hide()

				sendToGitButton.Show()
				sendToGitEntry.Show()
				GitProgressBar.Show()
				GitCheckMark.Show()

				ShowTimeDayLabel.Hide()
				ShowTimeUserLabel.Hide()
				ShowTimeEntry.Hide()
				ShowTimeRefresh.Hide()

				loadReadyDatesToEntry(readyToSendLog, sendToGitEntry)
				GetNumOfTasks()
			}
		}),
		// Python script
		widget.NewToolbarAction(theme.DocumentIcon(), func() {
			if !popActive {
				menu = "TimeLog"
				FYCT := yct.Format("2006-01-02")
				GetTimeLogs(FYCT, GetUsername(), "")
				if TimeLogs != nil {
					ShowTimeLogs()
				} else {
					ShowTimeEntry.Text = "No TimeLogs on date"
				}

				ShowTimeDayLabel.Text = FYCT
				ShowTimeUserLabel.Text = GetUsername()
				ShowTimeDayLabel.Refresh()
				ShowTimeUserLabel.Refresh()

				sidebar.Hide()
				selectEntry.Hide()
				summaryEntry.Hide()
				refeshButton.Hide()
				PlayButton.Hide()
				RecordButton.Hide()
				entry.Hide()
				image.Hide()

				PjLable.Hide()
				PjNavNeg.Hide()

				endTimeOP.Hide()
				timeOptionLabel.Hide()
				startTimeOP.Hide()
				timeMinusOP.Hide()
				saveOptions.Hide()
				ApiPopButton.Hide()
				WebButton.Hide()

				apiKeyLabel.Hide()
				apiKeyEntry.Hide()

				urlEntry.Hide()

				projectIDEntry.Hide()

				sendToGitButton.Hide()
				sendToGitEntry.Hide()
				GitProgressBar.Hide()
				GitCheckMark.Hide()

				ShowTimeDayLabel.Show()
				ShowTimeUserLabel.Show()
				ShowTimeEntry.Show()
				ShowTimeRefresh.Show()

				ShowTimeEntry.Refresh()
			}
		}),
		// OPTIONS MENU
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			menu = "Settings"
			sidebar.Hide()
			selectEntry.Hide()
			summaryEntry.Hide()
			refeshButton.Hide()
			PlayButton.Hide()
			RecordButton.Hide()
			entry.Show()
			image.Show()

			PjLable.Hide()

			endTimeOP.Show()
			timeOptionLabel.Show()
			startTimeOP.Show()
			timeMinusOP.Show()
			saveOptions.Show()
			ApiPopButton.Show()
			PjNavNeg.Show()
			WebButton.Show()

			apiKeyLabel.Show()
			apiKeyEntry.Show()

			urlEntry.Show()

			projectIDEntry.Show()

			sendToGitButton.Hide()
			sendToGitEntry.Hide()
			GitProgressBar.Hide()
			GitCheckMark.Hide()

			startTimeOP.Hide()
			timeMinusOP.Hide()

			dividers[2].Show()

			ShowTimeDayLabel.Hide()
			ShowTimeUserLabel.Hide()
			ShowTimeEntry.Hide()
			ShowTimeRefresh.Hide()

		}),
	)
}
