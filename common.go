package utils

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorIf(err error) bool {
	if err != nil {
		log.Print(err)
		return true
	}
	return false
}

func Log(data interface{}) {
	bytes, err := json.Marshal(data)
	ErrorIf(err)
	log.Println(string(bytes))
}

// 需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go语言的诞生时间 2006-01-02 15:04:05 -0700 MST。
func Date(t *time.Time) string {
	if t == nil {
		return "2006-01-02,15:04:05"
	}
	return t.Format("2006-01-02,15:04:05")
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
