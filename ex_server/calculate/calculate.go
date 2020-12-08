package calculate

import (
	"context"
	"fmt"
	"log"

	pb "github.com/KDJ0899/gRPCexample/ex_pb"
)

var (
	operators = map[pb.CalculateRequest_Operator]operator{
		pb.CalculateRequest_PLUS:  &plus{},
		pb.CalculateRequest_MINUS: &minus{},
		pb.CalculateRequest_MUL:   &multi{},
		pb.CalculateRequest_DIV:   &divide{},
	}
)

//Server calculate server
type Server struct{}

//Calculate calculate Numbers
func (s *Server) Calculate(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Printf("Numbers: %v  Operator %s", in.GetNumbers(), in.GetOperator().String())

	numbers := in.GetNumbers()
	op := operators[in.GetOperator()]

	result, err := op.calculate(numbers)

	if err != nil {
		return nil, err
	}

	return &pb.CalculateResponse{Result: result}, nil
}

type operator interface {
	calculate([]int64) (int64, error)
}

type plus struct{}
type minus struct{}
type multi struct{}
type divide struct{}

func (*plus) calculate(nums []int64) (int64, error) {
	var result int64 = 0
	for _, value := range nums {
		result += value
	}
	return result, nil
}

func (*minus) calculate(nums []int64) (int64, error) {
	var result int64 = 0
	for _, value := range nums {
		result -= value
	}
	return result, nil
}

func (*multi) calculate(nums []int64) (int64, error) {
	var result int64 = 0
	for _, value := range nums {
		result *= value
	}
	return result, nil
}

func (*divide) calculate(nums []int64) (int64, error) {
	var result int64 = 0
	for key, value := range nums {
		if key > 0 && value == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result += value
	}
	return result, nil
}
