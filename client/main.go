package main

import (
	"context"
	pb "grpc-demo/hello"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5050"
)

func main() {
	// 和grpc服务端建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "grpc"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("count not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
