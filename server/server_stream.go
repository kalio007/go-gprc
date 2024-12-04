package main

import (
	"log"
	"time"

	pb "github.com/kalio007/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NameList, stream pb.SimpleGreet_SayHelloServerStreamServer) error {
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
