package utils

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	_, exec := ExecCmdDir(strings.Join([]string{"date"}, " "), nil)
	exec(func(str string) {
		log.Println(str)
	})
}

func TestExecTop(t *testing.T) {
	cmd, exec := ExecCmdDir(strings.Join([]string{"top", "-pid 1"}, " "), nil)
	go func() {
		exec(func(str string) {
			log.Println(str)
		})
	}()

	time.Sleep(5 * time.Second)
	cmd.Process.Kill()
}

func TestExecSleep(t *testing.T) {
	cmd, exec := ExecCmdDir(strings.Join([]string{"sleep", "10"}, " "), nil)
	go func() {
		exec(func(str string) {
			log.Println(str)
		})
	}()

	time.Sleep(3 * time.Second)
	cmd.Process.Kill()
}
