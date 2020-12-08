package main

import (
	"context"
	"log"
	"os"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

const defaultName = "dongjin"

type hi struct{}

func (*hi) connectServer(ctx context.Context, conn *grpc.ClientConn) error {
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
