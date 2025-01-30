package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Melos")

	fileButton := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if f == nil {
				return
			}
			selectedFile := f.URI()
			fmt.Println("Selected file: ", selectedFile) // Do something with the selected file
		}, w)
	})

	w.SetContent(widget.NewLabel("Hello, Melos!"))
	w.SetContent(fileButton)
	w.ShowAndRun()

}
