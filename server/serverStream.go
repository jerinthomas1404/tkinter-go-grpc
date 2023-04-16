package main

import (
	"log"
	"time"

	pb "github.com/jerinthomas1404/tkinter-go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello from Server Streaming: " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}
