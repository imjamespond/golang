package stream

import (
	"codechiev/utils"
	"context"
	"io"
	"log"
	"test-grpc/test/stream/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:1808"
)

func RunClient(x int32) func() {
	return func() {
		conn, err := grpc.Dial(
			address,
			grpc.WithInsecure(), // 非安全
			grpc.WithBlock(),    // 阻塞
		)
		utils.FatalIf(err)
		defer conn.Close()
		cli := proto.NewStreamTestClient(conn)

		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			resp, err := cli.GetData(ctx, &proto.Point{X: 1, Y: 2})
			utils.ErrorIf(err)

			for {
				data, err := resp.Recv()
				if err == io.EOF {
					break
				}
				utils.FatalIf(err)
				log.Println(data)
			}
		}()
	}
}

func RunClient2(x int32) {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(), // 非安全
		grpc.WithBlock(),    // 阻塞
	)
	utils.FatalIf(err)
	defer conn.Close()
	cli := proto.NewStreamTestClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	stream, err := cli.SendNGetData(ctx)
	utils.FatalIf(err)
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			utils.FatalIf(err)
			log.Println(in)
		}
	}()
	for i := 0; i < 30; i++ {
		utils.FatalIf(stream.Send(&proto.Point{X: x, Y: int32(i)}))
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc
}
