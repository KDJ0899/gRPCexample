package main

import (
	"context"
	"log"
	"time"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "dongjin"

	plus   = pb.CalculateRequest_PLUS
	minus  = pb.CalculateRequest_MINUS
	multi  = pb.CalculateRequest_MUL
	divide = pb.CalculateRequest_DIV
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
