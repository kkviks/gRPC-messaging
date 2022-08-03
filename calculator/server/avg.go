package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"io"
	"log"
)

func (s *CalculateServer) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function in server called!")

	sum := int32(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count > 0 {
				average := float32(sum) / float32(count)
				return stream.SendAndClose(&pb.AvgResponse{Number: average})
			} else {
				return stream.SendAndClose(&pb.AvgResponse{Number: 0})
			}
		}
		if err != nil {
			log.Fatalf("err while reading client stream")
		}
		sum = sum + req.GetNumber()
		count++
	}
}
