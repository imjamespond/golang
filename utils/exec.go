package utils

import (
	"bufio"
	"context"
	"fmt"
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
func ExecCommand(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

type scan_func func(str string)
type exec_func func(scan scan_func)

func ExecCmdDir(command string, dir *string) (*exec.Cmd, exec_func) {
	parts := strings.Fields(command)
	return ExecCommandDir(dir, parts[0], parts[1:]...)
}
func ExecCommandDir(dir *string, name string, arg ...string) (*exec.Cmd, exec_func) {
	cmd := exec.Command(name, arg...)
	if dir != nil {
		cmd.Dir = *dir
	}
	return cmd, func(scan scan_func) {

		// out, err := cmd.Output()
		// PanicIf(err)

		stdout, err := cmd.StdoutPipe()
		PanicIf(err)
		stderr, err := cmd.StderrPipe()
		PanicIf(err)

		FatalIf(cmd.Start())
		fmt.Println("cmd.Process.Pid", cmd.Process.Pid)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		scanout := bufio.NewScanner(stdout)
		go func() {
			for scanout.Scan() {
				text := scanout.Text()
				if scan != nil {
					scan(text)
				}
			}
		}()

		scanerr := bufio.NewScanner(stderr)
		scanerr.Split(bufio.ScanRunes)
		go func() {
			for scanerr.Scan() {
				text := scanerr.Text()
				if scan != nil {
					scan(text)
				}
			}
		}()
		// slurp, err := io.ReadAll(stderr)
		// if err != nil || err != io.EOF {
		// 	fmt.Println(string("\033[31m"), string(slurp), string("\033[0m"))
		// }

		err = cmd.Wait()
		if err != nil {
			log.Println("Oops!", err.Error())
		}
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
