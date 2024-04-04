package main
  
import (   
	"context"
	"fmt"    

	pb "github.com/AdekunleDally/grpc-unary-api/calculator/proto"   
)

func (s *Server) AddTwoNumbers(ctx context.Context, numb *pb.AddTwoNumbRequest) (*pb.AdditionResponse, error) {
	switch numb.OperationType {
	case 1:
		return &pb.AdditionResponse{Result: numb.FirstNum + numb.SecondNum}, nil

	case 2:
		return &pb.AdditionResponse{Result: numb.FirstNum - numb.SecondNum}, nil
	case 3:
		return &pb.AdditionResponse{Result: numb.FirstNum * numb.SecondNum}, nil
	case 4:
		if numb.SecondNum != 0 {
			return &pb.AdditionResponse{Result: numb.FirstNum / numb.SecondNum}, nil
		} else {
			fmt.Println("Error: Division by zero")
			return &pb.AdditionResponse{ErrResult: "Error: Division by Zero"}, nil
		}
	default:
		fmt.Println("Invalid operation selected")
		return &pb.AdditionResponse{ErrResult: "Invalid operation selected"}, nil
	}
}
