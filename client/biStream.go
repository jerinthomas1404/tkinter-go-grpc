package main

import (
	"context"
	"io"
	"sync"
	"time"

	"log"

	pb "github.com/jerinthomas1404/tkinter-go-grpc/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("bidirectional streaming has started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				break
			}
			if err != nil {
				log.Fatalf("rror while streaming %v", err)
			}
			log.Println(message)
		}
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}

		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	wg.Wait()
	log.Printf("bidirectional streaming finished")
}
