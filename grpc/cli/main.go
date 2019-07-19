package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/myproto"
	"log"
	"os"
	"time"
)

const (
	address     = "127.0.0.1:8889"
	defaultName = "World!"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewGreetServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := cli.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
