package service_pdf

import (
	"4d-qrcode/util"
	"log"
	"strings"
)

func RunPdfkit(configPath string, rootDir string) {
	dir := "./pdfkit"
	cmd := []string{"node src/run.js", configPath, rootDir}
	util.ExecCmdDir(func(str string) {
		log.Println(str)
	})(strings.Join(cmd, " "), &dir)
}

func RunInstall() {
	dir := "./pdfkit"
	util.ExecCmdDir(func(str string) {
		log.Println(str)
	})(strings.Join([]string{"cnpm i"}, " "), &dir)
}
