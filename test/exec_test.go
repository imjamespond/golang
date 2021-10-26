package test

import (
	"codechiev/utils"
	"context"
	"log"
	"strings"
	"testing"
	"time"
)

func TestExecCmd(t *testing.T) {
	dir := "."
	cmd, exec := utils.ExecCmdDir(strings.Join([]string{"sleep", "10"}, " "), &dir)
	go func() {
		exec(func(str string) {
			log.Println(str)
		})
	}()

	time.Sleep(3 * time.Second)
	cmd.Process.Kill()
}

func TestExecCmdCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	utils.ExecCmdCtx(ctx)(strings.Join([]string{"sleep", "5"}, " "))
}

func TestExecCmdCtx1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	utils.ExecCmdCtx(ctx)(strings.Join([]string{"sleep", "50"}, " "))
}
