package main

import (
	"context"
	"log"

	pb "github.com/matizaj/oms/common/api"
	"google.golang.org/grpc"
)
type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server)  {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received!", p)
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}