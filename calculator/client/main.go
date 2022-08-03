package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting to %s: %v", addr, err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
	}(conn)

	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimeFactors(c)
	doAvg(c)
}
