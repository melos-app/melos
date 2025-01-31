package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func fileButton(w fyne.Window, callback func(string)) *widget.Button {
	return widget.NewButton("Select Typst file", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if f == nil {
				return
			}
			selectedFile := f.URI().String()
			callback(selectedFile)
		}, w)
	})
}

func main() {
	a := app.New()
	w := a.NewWindow("Melos")

	filePathDisplay := widget.NewLabel("")

	setFilePath := func(p string) {
		filePathDisplay.SetText(p)
	}

	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello, Melos!"),
		fileButton(w, setFilePath),
		filePathDisplay,
	))

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
