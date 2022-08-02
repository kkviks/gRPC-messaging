package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CalculateServer struct {
	pb.CalculatorServiceServer
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
	pb.RegisterCalculatorServiceServer(s, &CalculateServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
		return
	}

	log.Printf("Serve on %s\n", addr)
}
