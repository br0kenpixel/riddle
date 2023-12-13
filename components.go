package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var code *widget.Entry
var runtimeSelector *widget.Select
var output *widget.Entry

func MakeTitle() *canvas.Text {
	title := canvas.NewText("Riddle", color.White)
	title.TextStyle.Bold = true
	title.TextStyle.Monospace = true
	title.TextSize = 30.0

	return title
}

func MakeRuntimeSelection() *fyne.Container {
	label := widget.NewLabel("Runtime:")

	options := make([]string, 0)
	for i := range runtimes {
		options = append(options, fmt.Sprintf("%s [%s]", runtimes[i].path, runtimes[i].name))
	}

	selection := widget.NewSelect(options, func(_s string) {})
	selection.Selected = options[0]
	runtimeSelector = selection

	wrapper := container.New(layout.NewGridWrapLayout(fyne.NewSize(240, 35)), selection)

	hBox := container.NewHBox(label, wrapper)
	return hBox
}

func MakeInputSelection() *fyne.Container {
	label := widget.NewLabel("Input mode:")
	selection := widget.NewSelect([]string{"None", "Define...", "Serve file..."}, func(_s string) {})
	selection.PlaceHolder = "None"

	wrapper := container.New(layout.NewGridWrapLayout(fyne.NewSize(120, 35)), selection)

	hBox := container.NewHBox(label, wrapper)
	return hBox
}

func MakeCodeEntry() *widget.Entry {
	entry := widget.NewMultiLineEntry()
	entry.SetMinRowsVisible(15)
	entry.Text = "# Python code here...\nprint(\"Hello, World!\")"

	code = entry
	return entry
}

func MakeOutputConsole() *widget.Entry {
	entry := widget.NewMultiLineEntry()
	entry.SetMinRowsVisible(10)

	output = entry
	return entry
}

func MakeControls() *fyne.Container {
	runBtn := widget.NewButton("Run", RunCode)
	clearBtn := widget.NewButton("Clear console", ClearConsole)

	hBox := container.NewHBox(runBtn, clearBtn)
	return hBox
}
