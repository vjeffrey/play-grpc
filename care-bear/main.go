package main

import (
	"log"
	"net"

	"github.com/vjeffrey/play-grpc/unicorns-can-fly/unicorns-can-fly"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *unicorns.HelloRequest) (*unicorns.HelloReply, error) {
	return &unicorns.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *unicorns.HelloRequest) (*unicorns.HelloReply, error) {
	return &unicorns.HelloReply{Message: "Hello again " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	unicorns.RegisterFriendServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
