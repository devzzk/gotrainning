// 服务器端代码， 接收请求
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/devzzk/gotrainning/11grpc/helloword"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50001, "the server port")
)

// server is used to implement helloword.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloword.GreeterServer
func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.HelloReply{Message: "Hello: " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
