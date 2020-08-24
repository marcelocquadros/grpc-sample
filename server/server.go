package main

import (
	"context"
	"net"

	"github.com/marcelocquadros/grpc-sample/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	a, b := req.GetA(), req.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	a, b := req.GetA(), req.GetB()

	result := a * b

	return &pb.Response{Result: result}, nil
}

func main() {
	l, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalcServiceServer(grpcServer, &server{})

	if err = grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
