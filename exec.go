package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(command string) string {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

type scan_func func(str string)
type exec_func func(scan scan_func) string

func ExecCmdDir(command string, dir *string) (*exec.Cmd, exec_func) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	if dir != nil {
		cmd.Dir = *dir
	}
	return cmd, func(scan scan_func) string {

		// out, err := cmd.Output()
		// PanicIf(err)

		stdout, err := cmd.StdoutPipe()
		PanicIf(err)
		stderr, err := cmd.StderrPipe()
		PanicIf(err)

		FatalIf(cmd.Start())

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

		slurp, err := io.ReadAll(stderr)
		if err != nil || err != io.EOF {
			fmt.Println(string("\033[31m"), string(slurp), string("\033[0m"))
		}

		FatalIf(cmd.Wait())
		return out
	}
}

type exec_ctx_func func(command string)

func ExecCmdCtx(ctx context.Context) exec_ctx_func {
	return func(command string) {
		parts := strings.Fields(command)
		if err := exec.CommandContext(ctx, parts[0], parts[1:]...).Run(); err != nil {
			// will be interrupted.
			log.Println(err)
		}
	}
}
