package main

import (
	"gRPCexample/ex_server/calculate"
	"gRPCexample/ex_server/hi"
	"log"
	"net"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"

	Hi = iota
	Caculate
)

// server is used to implement helloworld.GreeterServer.

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	connectServer(s, Caculate)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func connectServer(s *grpc.Server, srvType int) {
	if srvType == Hi {
		pb.RegisterHiServer(s, &hi.Server{})
	} else if srvType == Caculate {
		pb.RegisterCalculateServer(s, &calculate.Server{})
	}
}
