package service_qrcode

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	qr "github.com/skip2/go-qrcode"
)

func TestGen1(t *testing.T) {
	err := qr.WriteColorFile("https://example.org", qr.High, 512, color.White, color.RGBA{255, 0, 0, 1}, "qr.png")
	if err != nil {
		panic(err)
	}
}

func TestGen2(t *testing.T) {

	file, _ := os.Create("qr.png")
	defer file.Close()

	q, err := qr.New("https://example.org", qr.Medium)
	if err != nil {
		panic(err)
	}

	q.DisableBorder = true
	q.ForegroundColor = color.RGBA{255, 255, 0, 255}
	q.BackgroundColor = color.White

	// err = q.WriteFile(256, "example-noborder.png")
	// if err != nil {
	// 	panic(err)
	// }

	img := q.Image(512)

	png.Encode(file, img)
}
