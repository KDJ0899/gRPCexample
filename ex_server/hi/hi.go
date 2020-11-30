package hi

import (
	"context"
	"log"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
)

//Server hi server
type Server struct{}

//SayHi test
func (s *Server) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HiReply{Messge: "Hi " + in.GetName()}, nil
}
