package main

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pool "github.com/processout/grpc-go-pool"
	"log"
	pb "igen.com/bee/protobuf/pbfile"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)


func newClient() (interface{}, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	p, err := pool.New(func() (*grpc.ClientConn, error) {
		return conn, nil
	}, 1, 5, 0)

	return p,err
}


func main() {
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	v, _ := newClient()
	for i := 0; i < 1000; i++ {
		go func() {
			pool := v.(*pool.Pool)
			myConn,_ := pool.Get(context.Background())
			client := pb.NewGreeterClient(myConn.ClientConn)

			name := "world"
			r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name:name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.Message)
			myConn.Close()
		}()
	}
	select {}


}
