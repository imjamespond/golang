package helloworld

import (
	"codechiev/utils"
	"context"
	"log"
	"test-grpc/test/hello/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:1808"
)

func RunClient() func(string) {
	return func(name string) {
		conn, err := grpc.Dial(
			address,
			grpc.WithInsecure(), // 非安全
			grpc.WithBlock(),    // 阻塞
		)
		utils.FatalIf(err)
		defer conn.Close()
		cli := proto.NewGreeterClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := cli.SayHello(ctx, &proto.HelloRequest{Name: name})
		utils.ErrorIf(err)
		log.Println(resp)
	}
}
