package service_qrcode

import (
	"testing"

	"4d-qrcode/model"
)

func TestJPEG(t *testing.T) {
	code := OpenJPEG("/Users/james/Downloads/codes/4-16220625519923o2kwj.jpg")
	tpl := OpenJPEG("/Users/james/Downloads/template.jpg")

	Process("./", &model.QRCodeConfig{})(tpl, code, "new.jpg")
}
