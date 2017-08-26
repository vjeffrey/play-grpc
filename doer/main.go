package main

import (
	"github.com/vjeffrey/play-grpc/unicorns-can-fly/unicorns-can-fly"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := unicorns.NewFriendClient(conn)

	// (context.Background(), &pb.HelloRequest{Name: name})
	c.StartChatApp(context.Background(), &unicorns.StartRequest{})
}
