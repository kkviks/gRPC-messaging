package main

import (
	pb "github.com/kkviks/grpc-messaging/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServer
}

const addr = "0.0.0.0:55051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s\n", addr)
		return
	}

	log.Printf("Listen on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve on %s\n", addr)
		return
	}

	log.Printf("Serve on %s\n", addr)
}
