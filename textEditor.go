package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor() {
	// a := app.New()
	// w := a.NewWindow("Text Editor")
	// w.Resize(fyne.NewSize(800, 600))

	w := myApp.NewWindow("Teax Editor")
	w.Resize(fyne.NewSize(500, 200))

	var count int = 1

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New_File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your Text....")
	input.Resize(fyne.NewSize(800, 600))

	saveBtn := widget.NewButton("Save File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName("New_File" + strconv.Itoa(count) + ".txt")

		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open File", func() {
		openFileDialg := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", readData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()
			}, w)

		openFileDialg.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}),
		)

		openFileDialg.Show()
	})

	textEditorContainer := container.NewVBox(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, textEditorContainer),
	)
	w.Show()
}
