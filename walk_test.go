package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Test1(t *testing.T) {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{Width: 600, Height: 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
}

func Test2(t *testing.T) {
	mw := new(MyMainWindow)

	var inTE, outTE *walk.TextEdit

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SCREAMO",
		MinSize:  Size{Width: 400, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &inTE},
			TextEdit{AssignTo: &outTE, ReadOnly: true},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					mw.openAction_Triggered()
				},
			},
		},
	}.Run()
}

func (mw *MyMainWindow) openAction_Triggered() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	fmt.Println(dlg.FilePath)

	return nil
}
