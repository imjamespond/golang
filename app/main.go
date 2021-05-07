package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/webview/webview"
)

func main() {

	fmt.Print(os.Args)

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	width, height := GetFullscreenSize()
	w.SetSize(width, height, webview.HintNone)
	w.SetTitle("Hello")

	_, release := os.LookupEnv("RELEASE")
	dev := flag.Bool("dev", false, "a bool")
	flag.Parse()
	url := "http://localhost:3000"
	if !*dev || release {
		url = "https://bing.com"
	}
	w.Navigate(url)
	w.Run()
}
