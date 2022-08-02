package main

import (
	"context"
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"io"
	"log"
)

func doPrimeFactors(c pb.CalculatorServiceClient) {
	log.Printf("doPrimeFactors is called!")

	req := &pb.PrimeRequest{Number: 210}

	stream, err := c.PrimeFactors(context.Background(), req)
	if err != nil {
		log.Printf("Error reciving prime factors: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading prime factors from stream: %v", err)
		}

		primeFactor := res.GetPrimeFactor()
		log.Printf("%v is one of the prime factors of %v\n", primeFactor, 210)
	}
}
