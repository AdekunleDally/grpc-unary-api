package main

import (  
	"log"
	"net" 
    
	  pb "github.com/AdekunleDally/grpc-unary-api/calculator/proto"
    
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8081"

type Server struct {
	pb.AdditionServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on :%s\n", addr)

	s := grpc.NewServer()
	pb.RegisterAdditionServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
