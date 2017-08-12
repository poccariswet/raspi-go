package main

import (
	"fmt"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)

	dialog := gtk.NewDialog()
	dialog.SetTitle("テスト用")
	vbox := dialog.GetVBox()
	label := gtk.NewLabel("Push down here")
	vbox.Add(label)

	button := gtk.NewButtonWithLabel("動画再生")
	button.Connect("clicked", func() {
		fmt.Println("動画を再生します...")
		gtk.MainQuit()
	})

	//button.SetSizeRequest(100, 100)
	vbox.Add(button)

	dialog.ShowAll()
	dialog.SetSizeRequest(200, 200)

	gtk.Main()
}
