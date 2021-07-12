package test

import (
	"4d-qrcode/util"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestCmd1(t *testing.T) {
	command := "pwd"
	parts := strings.Fields(command)
	data, err := exec.Command(parts[0], parts[1:]...).Output()
	if err != nil {
		panic(err)
	}

	output := string(data)
	log.Println(output)
}

func TestCmd2(t *testing.T) {
	dir := "../pdfkit"
	log.Println(util.ExecCmdDir(nil)("npm list", &dir))
}
