package main

import (
	"context"
	"log"

	pb "github.com/matizaj/oms/common/proto"
)

// type grpcHandler struct {
	
// }

func (s *server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("New order received %v\n", in)
	order := &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id: "1",
			CustomerId: "43",
			Status: "success",
		},
	}
	return order, nil
}

