package main

import (
	"context"
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg function in the client is called")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg")
	}
	for _, req := range reqs {
		log.Printf("Sending number to find avg: %v", req.Number)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response")
	}

	log.Printf("Average of %v is %v", reqs, res.GetNumber())
}
