package loadbalancing

import (
	"context"
	"log"
	"net"
	"sync"
	proto "test-grpc/test/hello/proto"

	"google.golang.org/grpc"
)

var (
	Addrs = []string{":50051", ":50052"}
)

type Server struct {
	proto.UnimplementedGreeterServer
	addr string
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "Hello " + req.Name + " from " + s.addr}, nil
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{addr: addr})
	log.Printf("serving on %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func StartServer() {
	var wg sync.WaitGroup
	for _, addr := range Addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}
