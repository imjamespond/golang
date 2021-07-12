package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"4d-qrcode/model"
	qr "4d-qrcode/service/qrcode"
	"4d-qrcode/util"
)

var cfg = util.ParseConfig("./config.json")
var qrcode = cfg["qrcode"].(map[string]interface{})
var codeX = int(qrcode["x"].(float64))
var codeY = int(qrcode["y"].(float64))
var width = int(qrcode["width"].(float64))
var heigth = int(qrcode["heigth"].(float64))

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
		cfg := model.QRCodeConfig{Width: width, Heigth: heigth, CodeX: codeX, CodeY: codeY}
		qr.Process(outputDir, &cfg)(
			qr.OpenJPEG(tpl),
			qr.OpenJPEG(filepath.Join(codeDir, file.Name())),
			filepath.Base(file.Name()))
	}

	fmt.Printf("总共用时：%f 秒", time.Since(before).Seconds())
}
