package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *CalculateServer) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max function called on the server side.")

	currMax := int32(math.MinInt32)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving client stream: %v", err)
		}

		// Consuming request
		currMax = int32(math.Max(float64(currMax), float64(req.GetNumber())))

		// Sending response
		err = stream.Send(&pb.MaxResponse{Number: currMax})
		if err != nil {
			log.Fatalf("Error sending response from server side: %v", err)
		}
	}
	return nil
}
