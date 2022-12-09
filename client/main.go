package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/danorel/gpb/client/proto"
)

const (
	defaultName = "world"
)

var (
	port = flag.Int("port", 9897, "The port to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type server struct {
	pb.UnimplementedChatServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	log.Printf("server %v listening at %v", name, lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}