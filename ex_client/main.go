package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "dongjin"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHiClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHi(ctx, &pb.HiRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessge())
}
