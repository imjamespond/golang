package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"4d-qrcode/model"
	pd "4d-qrcode/service/pdf"
	qr "4d-qrcode/service/qrcode"
	"4d-qrcode/util"
)

// 1,当前目录有config.json 2,传入template.jpg路径 3,template.jpg同目录有output目录

var linksPath = flag.String("links", "", "Download qrcode images from links.txt")
var nodeHome = flag.String("node_home", "./node", "node js home path")

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		log.Fatal("Please enter the template file!")
	}

	nodeHomePath, err := filepath.Abs(*nodeHome)
	util.PanicIf(err)
	os.Setenv("PATH", strings.Join([]string{os.Getenv("PATH"), nodeHomePath}, string(os.PathListSeparator)))
	// log.Println(os.Getenv("PATH"))

	cfgPath, err := filepath.Abs("./config.json")
	util.PanicIf(err)
	cfg := util.ParseConfig(cfgPath)
	qrcode := cfg["qrcode"].(map[string]interface{})

	rootDir, err := filepath.Abs(os.Args[1])
	util.PanicIf(err)
	tpl := filepath.Join(filepath.Dir(rootDir), "template.jpg")
	outputDir := filepath.Join(rootDir, "output")
	util.ErrorIf(os.Mkdir(outputDir, 0755))

	before := time.Now()

	// if (bool)(*gen) {
	if len(*linksPath) > 0 {
		links := qr.ReadLinks(rootDir)
		qrcodeCfg := model.GetQRCodeConfig(qrcode)
		tplImg := qr.OpenJPEG(tpl)

		for _, ln := range *links {
			log.Println(ln)

			img := qr.GetImage(ln)
			qr.Process(outputDir, qrcodeCfg)(tplImg, img, filepath.Base(ln))
		}
	}

	pd.RunPdfkit(cfgPath, rootDir)

	fmt.Printf("总共用时：%f 秒", time.Since(before).Seconds())
}
