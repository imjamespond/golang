package service_pdf

import (
	"4d-qrcode/util"
	"log"
	"strings"
)

func RunPdfkit(configPath string, rootDir string) {
	dir := "./pdfkit"
	cmd := []string{"node src/run.js", configPath, rootDir}
	log.Println(util.ExecCmdDir(func(str string) {
		log.Println(str)
	})(strings.Join(cmd, " "), &dir))
}
