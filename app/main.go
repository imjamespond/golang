package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/webview/webview"
)

func main() {

	fmt.Print(os.Args)

	dev := flag.Bool("dev", false, "a bool")
	flag.Parse()

	if !*dev {
		go StartWeb()
	}

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	width, height := GetFullscreenSize()
	w.SetSize(width, height, webview.HintNone)
	w.SetTitle("Hello")

	_, release := os.LookupEnv("RELEASE")
	url := "http://localhost:9000/app"
	if !*dev || release {
		url = "http://localhost:8080/app"
	}
	w.Navigate(url)
	w.Run()
}
