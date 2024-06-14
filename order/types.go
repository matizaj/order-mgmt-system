package main
import (
	"context"
	pb "github.com/matizaj/oms/common/api"
)

type OrderService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	CreateStore(context.Context) error
}