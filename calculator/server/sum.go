package main

import (
	"context"
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"log"
)

func (s *SumServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Value of request: %v", req)
	values := req.GetValues()
	sum := int32(0)
	for _, v := range values {
		sum += v
	}
	return &pb.SumResponse{Result: sum}, nil
}
