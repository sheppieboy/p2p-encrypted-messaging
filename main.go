package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	
	myApp := app.New()
	myWindow := myApp.NewWindow("Chat App")

	// Sidebar with a list of peers, current chat list, and chat requests list
	sidebar := container.NewVBox(
		widget.NewLabel("Peers List"),
		container.NewHBox(
			widget.NewLabel("Peer 1"),
			widget.NewButton("Request", func() {
				// Handle request button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Peer 2"),
			widget.NewButton("Request", func() {
				// Handle request button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Peer 3"),
			widget.NewButton("Request", func() {
				// Handle request button click
			}),
		),
		layout.NewSpacer(),
		widget.NewLabel("Current Chat List"),
		container.NewHBox(
			widget.NewLabel("Chat 1"),
			widget.NewButton("View", func() {
				// Handle view button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Chat 2"),
			widget.NewButton("View", func() {
				// Handle view button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Chat 3"),
			widget.NewButton("View", func() {
				// Handle view button click
			}),
		),
		layout.NewSpacer(),
		widget.NewLabel("Chat Requests List"),
		container.NewHBox(
			widget.NewLabel("Request 1"),
			widget.NewButton("Accept", func() {
				// Handle accept button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Request 2"),
			widget.NewButton("Accept", func() {
				// Handle accept button click
			}),
		),
		container.NewHBox(
			widget.NewLabel("Request 3"),
			widget.NewButton("Accept", func() {
				// Handle accept button click
			}),
		),
	)

	// Chat window
	chatWindow := container.NewVBox(
		widget.NewLabel("Chat Window"),
	)

	// Combine sidebar and chat window into a horizontal split container
	content := container.NewHSplit(sidebar, chatWindow)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
