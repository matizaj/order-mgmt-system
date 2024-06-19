package gateway

import (
	"context"

	pb "github.com/matizaj/oms/common/api"
)

type OrdersGateway interface {
	CreateOrder(ctx context.Context,  p *pb.CreateOrderRequest) (*pb.Order, error)
}