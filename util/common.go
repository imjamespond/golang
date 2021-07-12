package util

import (
	"bufio"
	"fmt"
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

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		cmd.Stdout = os.Stdout
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
		cmd.Wait()
		return out
	}
}
