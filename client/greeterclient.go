package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-go-example/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "server port")
	name = flag.String("name", "user", "user name")
)

func StartClient() {

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithInsecure())
	if err != nil {
		log.Fatal("error occured while connecting to grpc server")
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	greet_response, err := client.Greet(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatal("error occurred while invoking the greet server")
	}
	log.Printf("received response %v", greet_response.Greeting)
}

func main() {
	flag.Parse()
	StartClient()
}
