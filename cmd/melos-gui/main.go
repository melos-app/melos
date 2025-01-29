package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Melos")

	w.SetContent(widget.NewLabel("Hello, Melos!"))
	w.ShowAndRun()
}
