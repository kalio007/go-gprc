package main

import (
	"log"

	pb "github.com/kalio007/go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSimpleGreetClient(conn)

	// names := &pb.NameList{
	// 	Names: []string{"Kalio", "Princewill", "Riri"},
	// }

	callSayHello(client)
	//callSayHelloServerStream(client, names)
	// //callSayHelloClientStream(client, names)
	// callSayHelloBidirectionalStream(client, names)
}
