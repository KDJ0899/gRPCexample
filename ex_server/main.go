package main

import (
	"fmt"
	"gRPCexample/ex_server/calculate"
	"gRPCexample/ex_server/config"
	"gRPCexample/ex_server/hi"
	"log"
	"net"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
	"google.golang.org/grpc"
)

func main() {
	config.Init()

	lis, err := net.Listen("tcp", config.GetPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	if err := connectServer(s, config.GetServerType()); err != nil {
		log.Fatalf("failed connect server: %v", err)
	}

	log.Printf("Start %s Server", config.GetServerType())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func connectServer(s *grpc.Server, srvType string) error {
	if srvType == "hi" {
		pb.RegisterHiServer(s, &hi.Server{})
	} else if srvType == "calculate" {
		pb.RegisterCalculateServer(s, &calculate.Server{})
	} else {
		return fmt.Errorf("\"%s\" server is not exist", srvType)
	}
	return nil
}
