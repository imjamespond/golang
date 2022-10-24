package main

import (
	"os"
	"runtime"

	"github.com/xuri/excelize/v2"
)

func GetExcel(filePath string) *excelize.File {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
