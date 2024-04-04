package main

import (
	"log"  

	pb "github.com/AdekunleDally/grpc-unary-api/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:8081"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n ", err)
	}

	defer conn.Close()

	c := pb.NewAdditionServiceClient(conn)

	doAddTwoNumbers(c)
}
