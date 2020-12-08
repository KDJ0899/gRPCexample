package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := connectCaculateServer(ctx, conn); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
