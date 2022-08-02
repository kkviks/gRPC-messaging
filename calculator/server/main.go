package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SumServer struct {
	pb.SumServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen to p: %v\n", addr)
		return
	}
	log.Printf("Listen on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &SumServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		return
	}

	log.Printf("Serve on %s\n", addr)
}
