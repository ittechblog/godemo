package main

import (
	"log"
	pb "igen.com/bee/grpc"
	"os"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	address     = "localhost:9998"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSimpleClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1];
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name:name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	conn.Close()
	log.Printf("Greeting: %s", r.Message)

}
