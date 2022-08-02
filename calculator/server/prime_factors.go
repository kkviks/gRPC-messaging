package main

import (
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"log"
)

func (s *CalculateServer) PrimeFactors(req *pb.PrimeRequest, stream pb.CalculatorService_PrimeFactorsServer) error {
	log.Printf("PrimeFactors called with: %v", req)

	k := uint32(2)
	N := req.GetNumber()

	for N > 1 {
		if N%k == 0 {
			if err := stream.Send(&pb.PrimeResponse{PrimeFactor: k}); err != nil {
				log.Printf("error sending prime factor: %v", err)
				return err
			}
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
