package main

import (
	"4d-qrcode/util"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

var cfg = util.ParseConfig("./config.json")
var qrcode = cfg["qrcode"].(map[string]interface{})
var codeX = int(qrcode["x"].(float64))
var codeY = int(qrcode["y"].(float64))

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please enter the template file!")
	}

	tpl := os.Args[1]
	codeDir := filepath.Join(filepath.Dir(tpl), "codes")
	outputDir := filepath.Join(filepath.Dir(tpl), "output")

	codeFiles, err := ioutil.ReadDir(codeDir)
	if err != nil {
		log.Fatal(err)
	}

	before := time.Now()

	for _, file := range codeFiles {
		if file.IsDir() {
			continue
		}
		ext := filepath.Ext(file.Name())
		ext = strings.ToLower(ext)
		if ext != ".jpg" {
			continue
		}

		// fmt.Println(file.Name(), file.IsDir(), ext)
		log.Println(file.Name())
		process(outputDir)(
			openJPEG(tpl),
			openJPEG(filepath.Join(codeDir, file.Name())),
			filepath.Base(file.Name()))
	}

	fmt.Printf("总共用时：%f 秒", time.Since(before).Seconds())
}

func openJPEG(file string) *image.Image {
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

func process(outputDir string) process_func {
	return func(_tpl *image.Image, _code *image.Image, codeName string) {
		code := resize.Resize(80, 80, *_code, resize.Lanczos3)
		tpl := resize.Resize(150, 0, *_tpl, resize.Lanczos3) //preserve aspect ratio

		m := image.NewRGBA(tpl.Bounds())
		draw.Draw(m, m.Bounds(), tpl, image.Point{0, 0}, draw.Src)
		draw.Draw(m, m.Bounds(), code, image.Point{codeX, codeY}, draw.Src)

		toimg, _ := os.Create(filepath.Join(outputDir, codeName))
		defer toimg.Close()

		jpeg.Encode(toimg, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
}
