package main

import (
	"context"

	pb "github.com/jerinthomas1404/tkinter-go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello from unary",
	}, nil

}
