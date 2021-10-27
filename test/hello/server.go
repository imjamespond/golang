package helloworld

import (
	"codechiev/utils"
	"context"
	"net"

	"test-grpc/test/hello/proto"

	"google.golang.org/grpc"
)

const (
	port = ":1808"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "Hello " + req.Name}, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", port)
	utils.FatalIf(err)

	srv := grpc.NewServer()
	proto.RegisterGreeterServer(srv, &server{})

	err = srv.Serve(lis)
	utils.FatalIf(err)
}
