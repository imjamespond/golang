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

var links = flag.String("links", "", "Download qrcode images from links.txt")
var nodeHome = flag.String("node_home", "./node", "node js home path")
var install = flag.Bool("install", false, "run npm install")

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		log.Fatal("Please enter the template file!")
	}

	nodeHomePath, err := filepath.Abs(*nodeHome)
	util.FatalIf(err)
	os.Setenv("PATH", strings.Join([]string{os.Getenv("PATH"), nodeHomePath}, string(os.PathListSeparator)))
	// log.Println(os.Getenv("PATH"))

	cfgPath, err := filepath.Abs("./config.json")
	util.FatalIf(err)
	cfg := util.ParseConfig(cfgPath)
	qrcode := cfg["qrcode"].(map[string]interface{})

	args1 := os.Args[len(os.Args)-1]
	var rootDir string
	if isDir, err := util.IsDirectory(args1); !isDir {
		util.FatalIf(err)
		log.Println(args1, "isDir", isDir)
		rootDir = filepath.Dir(args1)
	} else {
		rootDir, err = filepath.Abs(args1)
		util.FatalIf(err)
	}
	tpl := filepath.Join(rootDir, "template.jpg")
	outputDir := filepath.Join(rootDir, "output")
	util.ErrorIf(os.Mkdir(outputDir, 0755))

	before := time.Now()

	if (bool)(*install) {
		pd.RunInstall()
	}

	if len(*links) > 0 {
		linksFile, err := filepath.Abs(*links)
		util.FatalIf(err)
		links := qr.ReadLinks(linksFile)
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
