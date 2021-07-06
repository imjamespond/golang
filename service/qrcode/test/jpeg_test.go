package main

import (
	"testing"

	qr "4d-qrcode/service/qrcode"
)

func TestJPEG(t *testing.T) {
	code := qr.OpenJPEG("/Users/james/Downloads/codes/4-16220625519923o2kwj.jpg")
	tpl := qr.OpenJPEG("/Users/james/Downloads/template.jpg")

	qr.Process("./")(tpl, code, "new.jpg")
}
