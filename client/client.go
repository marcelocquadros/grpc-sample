package main

import (
	"context"
	"fmt"

	"github.com/marcelocquadros/grpc-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial(":4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer con.Close()

	svc := pb.NewCalcServiceClient(con)

	req := &pb.Request{A: 3, B: 9}

	res, err := svc.Sum(context.Background(), req)

	if err != nil {
		fmt.Println("Error caling server")
		return
	}

	fmt.Printf("Result of Sum is: %v \n", res.Result)

	res, err = svc.Multiply(context.Background(), req)

	if err != nil {
		fmt.Println("Error caling server")
		return
	}

	fmt.Printf("Result of Multiply is: %v \n", res.Result)
}
