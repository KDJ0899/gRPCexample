package main

import (
	"context"
	"gRPCexample/ex_client/config"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	clients = map[string]client{
		"hi":        &hi{},
		"calculate": &calculate{},
	}
)

type client interface {
	connectServer(context.Context, *grpc.ClientConn) error
}

func main() {
	config.Init()
	address := config.GetServerAddress() + ":" + config.GetPort()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Try Connect %s Server", config.GetClientType())
	if err := clients[config.GetClientType()].connectServer(ctx, conn); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
