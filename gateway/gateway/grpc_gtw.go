package gateway

import (
	"context"
	"fmt"
	"log"

	"github.com/matizaj/oms/common/discovery"
	pb "github.com/matizaj/oms/common/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gateway struct {
	registry discovery.Registry
}

func NewGrpcGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}

func (g *gateway) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// TODO: create fun to reuse for getting grpc connection
	gRPConn , err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewOrderServiceClient(gRPConn)

	order :=  &pb.CreateOrderRequest{
		CustomerId: in.CustomerId,
		Items: in.Items,
	}
	o, err := grpcClient.CreateOrder(ctx,order)
	if err != nil {
		log.Printf("failed to create order %v\n", err)
		errGrpc := status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed to create order %v", err),
		)
		return nil, errGrpc
	}

	return o, nil
}

func (g *gateway) GetOrder(ctx context.Context, customerId, orderId string) (*pb.GetOrderResponse, error) {
	log.Printf("client invoke get order")

	gRPConn , err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewOrderServiceClient(gRPConn)
	log.Println(grpcClient)

	req := &pb.GetOrderRequest{
		CustomerId: customerId,
		OrderId: orderId,
	}
	result, err := grpcClient.GetOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}