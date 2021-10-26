package test

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"sd-2110/proto/hello"

	"google.golang.org/grpc"
	"my.com/utils"
)

const (
	address = "localhost:50051"
)

func TestHello(t *testing.T) {
	dir := "../java/target/"
	_, exec := utils.ExecCmdDir(strings.Join([]string{"java", "--version"}, " "), &dir)
	go func() {
		exec(func(str string) {
			log.Println(str)
		})
	}()
	RunClient()("Foobar")
}

func RunClient() func(string) {
	return func(name string) {
		conn, err := grpc.Dial(
			address,
			grpc.WithInsecure(), // 非安全
			grpc.WithBlock(),    // 阻塞
		)
		utils.FatalIf(err)
		defer conn.Close()
		cli := hello.NewGreeterClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := cli.SayHello(ctx, &hello.HelloRequest{Name: name})
		utils.ErrorIf(err)
		log.Println(resp)
	}
}
