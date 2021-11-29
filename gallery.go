package main

import (
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)

func showGalleryApp(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(800, 600))
	root_src := "C:\\Users\\ASUS\\Pictures\\Saved Pictures"

	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	tabs := container.NewAppTabs()
	for _, file := range files {
		if !file.IsDir() {
			extension_Array := strings.Split(file.Name(), ".")
			extension := extension_Array[len(extension_Array)-1]
			if extension == "png" || extension == "jpeg" || extension == "jpg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	galleryContainer := tabs

	// w.SetContent(tabs)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, galleryContainer))
	w.Show()
}
