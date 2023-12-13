package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var runtimes []Runtime

func main() {
	a := app.New()
	w := a.NewWindow("Riddle")
	w.Resize(fyne.NewSize(800, 100))

	runtimes = GetRuntimes()

	if len(runtimes) == 0 {
		dialog.NewInformation("Riddle", "No Python installations found", w).Show()
		return
	}

	vBox := container.NewVBox(MakeTitle(), MakeRuntimeSelection(), MakeCodeEntry(), widget.NewSeparator(), MakeInputSelection(), MakeOutputConsole(), MakeControls())

	w.SetContent(vBox)
	w.ShowAndRun()
}
