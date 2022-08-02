package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = "0.0.0.0:55051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s\n", addr)
		return
	}

	log.Printf("Listen on %s\n", addr)

	s := grpc.NewServer()

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve on %s\n", addr)
		return
	}

	log.Printf("Serve on %s\n", addr)
}
