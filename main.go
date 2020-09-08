package main

import (
	"context"
	pb "grpc-demo/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

// 来实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
