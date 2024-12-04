package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kalio007/go-grpc/proto"
)

func callSayHelloServerStream(client pb.SimpleGreetClient, names *pb.NameList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}

	log.Printf("Streaming finished")
}
