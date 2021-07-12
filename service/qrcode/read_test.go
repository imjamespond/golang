package service_qrcode

import (
	"4d-qrcode/model"
	"4d-qrcode/util"
	"bufio"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var file = os.Getenv("HOME") + "/Downloads/2021429104441-1.txt"
var tpl = os.Getenv("HOME") + "/Downloads/template.jpg"

func TestReadText(t *testing.T) {
	log.Println(file)

	pfile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer pfile.Close()

	scanner := bufio.NewScanner(pfile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		log.Println(filepath.Base(scanner.Text()))

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestReadImg(t *testing.T) {
	resp, err := http.Get("https://ymvideos.oss-cn-shanghai.aliyuncs.com/Production/0/2021429104441-1/1-1619664281976szj-sa.png")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	g := img.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	// The resolution is height x width
	resolution := height * width

	// Print results
	log.Println(resolution, "pixels")
}

var cfg = util.ParseConfig("../../config.json")
var qrcode = cfg["qrcode"].(map[string]interface{})

// var codeX = int(qrcode["x"].(float64))
// var codeY = int(qrcode["y"].(float64))
// var width = int(qrcode["width"].(float64))
// var heigth = int(qrcode["heigth"].(float64))

func TestReadNCompose(t *testing.T) {
	links := ReadLinks(file)
	qrcodeCfg := model.GetQRCodeConfig(qrcode)
	tplImg := OpenJPEG(tpl)

	for _, ln := range *links {
		log.Println(ln)

		img := GetImage(ln)
		// cfg := model.QRCodeConfig{Width: width, Heigth: heigth, CodeX: codeX, CodeY: codeY}
		Process(os.Getenv("HOME")+"/Downloads/output", qrcodeCfg)(tplImg, img, filepath.Base(ln)) //strconv.Itoa(i)+".jpg")
	}
}
