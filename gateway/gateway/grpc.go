package gateway

import (
	"context"

	pb "github.com/matizaj/oms/common/api"
	"github.com/matizaj/oms/common/discovery"
)

type gateway struct {
	registry discovery.Registry 
}


func NewGrpcGateway(registry discovery.Registry) *gateway{
	return &gateway{registry}
}

func (g *gateway) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)	
	if err != nil {
		return nil, err
	}

	c :=pb.NewOrderServiceClient(conn)

	order, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerId: p.CustomerId,
		Items: p.Items,
	})

	return order, nil
}