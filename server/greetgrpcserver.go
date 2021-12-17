package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-go-example/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const PROTOCOL = "tcp"

var (
	port = flag.Int("port", 50051, "server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("received request %v", request.Name)
	return &pb.HelloResponse{Greeting: "Hello " + request.Name}, nil
}

func StartGrpCServer() {
	listener, error := net.Listen(PROTOCOL, fmt.Sprintf("localhost:%d", *port))
	if error != nil {
		log.Fatalf("failed to start the server %v", error)
	}
	log.Printf("started server on port %v", *port)
	grpc_server := grpc.NewServer()
	pb.RegisterGreeterServer(grpc_server, &server{})
	if error := grpc_server.Serve(listener); error != nil {
		log.Fatalf("failed to serve %v", error)
	}
}
func main() {
	flag.Parse()
	StartGrpCServer()
}
