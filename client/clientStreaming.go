package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jerinthomas1404/tkinter-go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming has started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error whle sending %v", err)
		}
		log.Printf("sent the request with the name :%v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("streaming finished")
	if err != nil {
		log.Fatalf("error while closing %v", err)
	}

	log.Printf("%v", res.Messages)
}
