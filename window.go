package main

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func setWin(win fyne.Window) { // since window is interface, no need to pass pointer anymore
	cfg := GetCfg()

	// input := widget.NewEntry()
	// input.SetPlaceHolder("Enter url")
	textAreaUrl := widget.NewMultiLineEntry()
	textAreaUrl.SetPlaceHolder("Enter url")
	textAreaUrl.Wrapping = fyne.TextWrapWord

	bashLabel := widget.NewLabel(getBash(cfg.Bash))

	bashBtn := widget.NewButton("Bash", func() {
		file := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
			if e == nil && uc != nil {
				log.Println(uc.URI().Path())
				cfg.Bash = uc.URI().Path()
				WriteCfg()
				bashLabel.Text = getBash(cfg.Bash)
			}
		}, win)
		file.Show()
	})

	reqBtn := widget.NewButton("请求", func() {
		// log.Println("Content was:", textAreaUrl.Text)
		if len(cfg.Bash) == 0 {
			dialog.NewError(fmt.Errorf("please set bash location"), win).Show()
		} else if len(textAreaUrl.Text) > 0 && strings.Index(textAreaUrl.Text, "curl") == 0 {
			curl(cfg.Bash, textAreaUrl.Text)
		}
	})

	row := container.New(layout.NewHBoxLayout(), bashBtn, reqBtn)

	content := container.NewVBox(textAreaUrl, bashLabel, row)

	win.SetContent(content)
}

func getBash(path string) string {
	return "bash location: " + path
}
