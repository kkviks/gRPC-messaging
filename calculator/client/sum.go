package main

import (
	"context"
	pb "github.com/kkviks/grpc-messaging/calculator/proto"
	"log"
)

func doSum(c pb.SumServiceClient) {
	log.Printf("doSum is called!")

	ctx := context.Background()
	req := &pb.SumRequest{
		Values: []int32{1, 2, 3},
	}
	res, err := c.Sum(ctx, req)
	if err != nil {
		log.Printf("Error in finding sum of %v: %v", req.GetValues(), err)
	}

	log.Printf("The sum of %v is: %v", req.GetValues(), res.GetResult())
}
