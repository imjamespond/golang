package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func Curl(command string) string {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = strings.Replace(command, "curl", "curl.exe -sS", 1)
		// cmd = strings.Replace(cmd, "\"", "", -1)
	} else {
		cmd = strings.Replace(command, "curl", "curl -sS", 1)
	}
	log.Println(cmd)
	execCmd := (*exec.Cmd)(nil)
	if runtime.GOOS == "windows" {
		execCmd = exec.Command("cmd", "/c", cmd)
	} else {
		execCmd = exec.Command("bash", "-c", cmd)
	}

	out, err := execCmd.CombinedOutput()
	// log.Println(string(out))
	if err != nil {
		log.Println(err)
	}
	return string(out)
}

func getPath(file string) string {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "darwin" {

		// dir, err = os.Executable()
		dir, err = os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		dir = filepath.Join(dir, "Downloads", "swagger2ts")
		// dir = strings.TrimSuffix(dir, "/swagger2ts.app/Contents/MacOS/swagger2ts")
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}

	path := filepath.Join(dir, file)
	log.Println(path)

	return path
}

func Save(file string, val string) {

	path := getPath(file)

	err := os.WriteFile(path, []byte(val), 0666)
	if err != nil {
		log.Println(err)
	}

}

func Read(file string) *string {

	path := getPath(file)

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	str := string(data)

	return &str
}
