package main

import (
	"context"
	pb "github.com/matizaj/oms/common/proto"
)

type OrderService interface {
	 CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
	 ValidateOrder(context.Context, *pb.CreateOrderRequest) error	
	 GetOrder(context.Context, *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
}

type OrderStore interface {
	Create(context.Context) error
	Get(context.Context) error
}