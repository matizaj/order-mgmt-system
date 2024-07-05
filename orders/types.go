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
	Create(ctx context.Context, order Order) error
	Get(ctx context.Context, customerId, orderId string) *Order
}

type Order struct {
	Order *pb.Order
}