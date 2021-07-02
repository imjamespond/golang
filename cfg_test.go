package main

import (
	"log"
	"testing"

	"4d-qrcode/util"
)

func TestJSON(t *testing.T) {
	cfg := util.ParseConfig("./config.json")
	qrcode := cfg["qrcode"].(map[string]interface{})
	log.Println(qrcode["x"], qrcode["y"])
}
