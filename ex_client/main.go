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

func connectHiServer(ctx context.Context, conn *grpc.ClientConn) error {
	c := pb.NewHiClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHi(ctx, &pb.HiRequest{Name: name})
	if err != nil {
		return err
	}

	log.Printf("Response: %s", r.GetMessge())

	return nil
}

func connectCaculateServer(ctx context.Context, conn *grpc.ClientConn) error {
	c := pb.NewCalculateClient(conn)
	numbers := []int64{1, 2, 4}

	r, err := c.Calculate(ctx, &pb.CalculateRequest{Numbers: numbers, Operator: plus})
	if err != nil {
		return err
	}

	log.Printf("Result: %d", r.Result)
	return nil
}
