package stream

import (
	"codechiev/utils"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"test-grpc/test/stream/proto"

	"google.golang.org/grpc"
)

const (
	port = ":1808"
)

type server struct {
	proto.UnimplementedStreamTestServer
}

func (s *server) GetData(req *proto.Point, stream proto.StreamTest_GetDataServer) error {
	for i := 0; i < 3; i++ {
		stream.Send(&proto.Response{Location: req, Name: "Server 1"})
		time.Sleep(time.Second >> 1)
	}
	return nil
}

func (s *server) SendNGetData(stream proto.StreamTest_SendNGetDataServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("client closed")
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(in)
		if err := stream.Send(&proto.Response{Location: in, Name: "Foobar"}); err != nil {
			return err
		}

		// auto close client?
		if r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31(); r%30 == 1 {
			log.Println("rand is", r)
			return nil
		} else {
			log.Println("wait...", r)
		}
	}
}

func RunServer() {
	lis, err := net.Listen("tcp", port)
	utils.FatalIf(err)

	srv := grpc.NewServer()
	proto.RegisterStreamTestServer(srv, &server{})

	err = srv.Serve(lis)
	utils.FatalIf(err)
}
