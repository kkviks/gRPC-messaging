package main

import (
	"context"
	pb "github.com/kkviks/grpc-messaging/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet is called!\n")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Vikas"})
	if err != nil {
		log.Fatalf("Couldn't get greeting %v\n", err)
		return
	}

	log.Printf("Response from Server: %v\n", res.Result)
}
