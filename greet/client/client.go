package main

import (
	pb "github.com/kkviks/grpc-messaging/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect on %s\n", addr)
	}

	log.Printf("Connected to %s\n", addr)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", conn)
		}
		log.Println("Connection closed")
	}(conn)

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
