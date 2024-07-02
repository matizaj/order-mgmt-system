package gateway

import (
	"context"

	pb "github.com/matizaj/oms/common/proto"
)

type OrderGateway interface {
	CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}