package main

import (
	"context"
	"fmt"
	"log"  

	pb "github.com/AdekunleDally/grpc-unary-api/calculator/proto"  
)
  
type Operation struct {
	OperationType string
	FirstNum      float64
	SecondNum     float64
	Result        float64
}

func doAddTwoNumbers(c pb.AdditionServiceClient) {
	var operationType int64
	fmt.Println("Welcome to the Calculator App. Please Select an Operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Println("Enter the number corresponding to the operation: ")
	fmt.Scanln(&operationType)

	if operationType > 4 {
		fmt.Println("You have selected an Invalid Option. Please Enter a valid option")
		return
	}  

	var first, second float64
	fmt.Println("Please enter the first Number:")
	fmt.Scanln(&first)

	fmt.Println("Please enter the second Number:")
	fmt.Scanln(&second)
	if second == 0 {
		fmt.Println("Error: Division by Zero")
		return
	}

	res, err := c.AddTwoNumbers(context.Background(), &pb.AddTwoNumbRequest{
		FirstNum:      first,
		SecondNum:     second,
		OperationType: operationType,
	})
	if err != nil {
		log.Fatalf("Could not carry out the computation: %v\n", err)
	}

	log.Printf("The result is  : %v\n", res.Result)
}
