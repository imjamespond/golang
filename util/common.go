package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
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

func ErrorIf(err error) {
	if err != nil {
		log.Print(err)
	}
}

func ExecCmd(command string) string {
	return ExecCmdDir(nil)(command, nil)
}

type scan_func func(str string)
type exec_func func(command string, dir *string) string

func ExecCmdDir(scan scan_func) exec_func {
	return func(command string, dir *string) string {
		parts := strings.Fields(command)
		cmd := exec.Command(parts[0], parts[1:]...)
		if dir != nil {
			cmd.Dir = *dir
		}

		// out, err := cmd.Output()
		// PanicIf(err)

		stdout, err := cmd.StdoutPipe()
		PanicIf(err)
		stderr, err := cmd.StderrPipe()
		PanicIf(err)
		cmd.Start()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		scanner := bufio.NewScanner(stdout)
		var out string
		for scanner.Scan() {
			m := scanner.Text()
			if scan == nil {
				fmt.Println(m)
				out += m
			} else {
				scan(m)
			}
		}

		slurp, _ := io.ReadAll(stderr)
		if scan == nil {
			fmt.Println(string(slurp))
		} else {
			scan(string(slurp))
		}

		FatalIf(cmd.Wait())
		return out
	}
}

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
