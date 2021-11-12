package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func init() {
	os.Setenv("FYNE_FONT", "c:\\WINDOWS\\Fonts\\MSYH.TTC ")
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800, 600))

	setWin(w)

	w.ShowAndRun()
}
