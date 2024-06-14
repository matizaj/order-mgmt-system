package main

import (
	"context"
	"log"

	pb "github.com/matizaj/oms/common/api"
	"google.golang.org/grpc"
)
type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	orderService OrderService
}

func NewGrpcHandler(grpcServer *grpc.Server, oService OrderService)  {
	handler := &grpcHandler{orderService: oService}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received!", p)
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}