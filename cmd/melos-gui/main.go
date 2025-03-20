package main

import (
	"fmt"

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

	// steps:
	//  - uncompress
	//  - removeTitles
	//  - generateSvg
	//  - updateTypst
	//  - makeBook

	uncompress := false
	removeTitles := false
	generateSvg := false
	updateTypst := false
	makeBook := false

	// TODO: set checkbox to checked by default

	w.SetContent(container.NewVBox(
		widget.NewLabel("Hello, Melos!"),
		fileButton(w, setFilePath),
		filePathDisplay,
		widget.NewCheck("Uncompress", func(checked bool) {
			uncompress = checked
		}),
		widget.NewCheck("Remove Titles", func(checked bool) {
			removeTitles = checked
		}),
		widget.NewCheck("Generate SVG", func(checked bool) {
			generateSvg = checked
		}),
		widget.NewCheck("Update Typst file", func(checked bool) {
			updateTypst = checked
		}),
		widget.NewCheck("Make book", func(checked bool) {
			makeBook = checked
		}),
		widget.NewButton("Run", func() {
			fmt.Println("uncompress:", uncompress)
			fmt.Println("removeTitles:", removeTitles)
			fmt.Println("generateSvg:", generateSvg)
			fmt.Println("updateTypst:", updateTypst)
			fmt.Println("makeBook:", makeBook)
		}),
	))

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
