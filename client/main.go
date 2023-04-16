package main

import (
	"log"

	pb "github.com/jerinthomas1404/tkinter-go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
}
