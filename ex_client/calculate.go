package main

import (
	"context"
	"log"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

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
