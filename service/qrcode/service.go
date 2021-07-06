package service_qrcode

import (
	"4d-qrcode/util"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

var cfg = util.ParseConfig("./config.json")
var qrcode = cfg["qrcode"].(map[string]interface{})
var codeX = int(qrcode["x"].(float64))
var codeY = int(qrcode["y"].(float64))
var width = int(qrcode["width"].(float64))
var heigth = int(qrcode["heigth"].(float64))

func OpenJPEG(file string) *image.Image {
	// Decode the JPEG data. If reading from file, create a reader with
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return &m
}

type process_func func(_tpl *image.Image, _code *image.Image, codeName string)

func Process(outputDir string) process_func {
	return func(_tpl *image.Image, _code *image.Image, codeName string) {
		code := resize.Resize(uint(width), uint(heigth), *_code, resize.Lanczos3)
		// tpl := resize.Resize(150, 0, *_tpl, resize.Lanczos3) //preserve aspect ratio
		tpl := *_tpl

		m := image.NewRGBA(tpl.Bounds())
		draw.Draw(m, m.Bounds(), tpl, image.Point{0, 0}, draw.Src)
		draw.Draw(m, m.Bounds(), code, image.Point{codeX, codeY}, draw.Src)

		toimg, _ := os.Create(filepath.Join(outputDir, codeName))
		defer toimg.Close()

		jpeg.Encode(toimg, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
}
