package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/danorel/gpb/client/chat"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:9898", "The address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type client struct {
	pb.UnimplementedChatServer
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	log.Printf("client %v listening at %v", *name, *addr)

	client := pb.NewChatClient(conn)
	reply, err := client.SendMessage(context.Background(), &pb.Message{
		Id:    0,
		Title: "Hello, server!",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Message: %s", reply.Title)
}