package main

import (
	"fmt"
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

func (s *server) StartChatApp(ctx context.Context, in *unicorns.StartRequest) (*unicorns.Response, error) {

	var response int
	fmt.Println("This program will start a chat app, are you sure about this?\n")

	fmt.Scanf("%c", &response) //<--- here
	switch response {
	default:
		fmt.Println("Chat launch aborted!")
	case 'y':
		fmt.Println("Chat launched")
	case 'Y':
		fmt.Println("Chat launched")
	}
	return &unicorns.Response{}, nil
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
