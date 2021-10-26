package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"my.com/utils"
)

func main() {
	purify()

	dir := "./java/target/"
	cmd, exec := utils.ExecCmdDir(
		strings.Join([]string{"java", "-cp sd-2110-1.0-SNAPSHOT-jar-with-dependencies.jar", "sd.hello.App", "args1", "args2"}, " "),
		&dir)
	go func() {
		exec(func(str string) {
			log.Println(str)
		})
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	s := <-sigchan
	log.Println(s)
	cmd.Process.Kill()
}

func purify() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	out := utils.ExecCmd("jps")
	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)
		if cols[1] == "App" {
			pid, err := strconv.ParseInt(cols[0], 10, 0)
			utils.PanicIf(err)
			p, err := os.FindProcess(int(pid))
			utils.PanicIf(err)
			p.Kill()
		}
	}
}
