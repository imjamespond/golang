package main

import (
	"testing"
)

func TestJPEG(t *testing.T) {
	code := openJPEG("/Users/james/Downloads/codes/4-16220625519923o2kwj.jpg")
	tpl := openJPEG("/Users/james/Downloads/template.jpg")

	process("./")(tpl, code, "new.jpg")
}
