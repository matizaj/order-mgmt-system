package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/matizaj/oms/common/proto"
)

type service struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &pb.CreateOrderResponse{
		Order: &pb.Order{
			CustomerId: "1",
			Status: "success",
			Items: []*pb.Item{
				{
					Id: "1",
					Name: "rope",
					Quantity: 1,
					PriceId: "price_1PYSN0EJEwxXWrvp5iF9aTfD",
				},
				
			},
		},
	}
	orderToStore := Order{
		Order: order.Order,
		}
	
	
	err := s.store.Create(ctx, orderToStore)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *service) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order := &pb.GetOrderResponse{
		Order: &pb.Order{
			Id: in.OrderId,
			CustomerId: in.CustomerId,
			Status: "success",
			Items: []*pb.Item{
				{
					Id: "1",
					Name: "rope",
					Quantity: 5,
					PriceId: "1111--2222-333",
				},
			},
		},
	}

	return order, nil
}

func (s *service) ValidateOrder(ctx context.Context, in *pb.CreateOrderRequest) error	 {
	if len(in.Items) <= 0 {
		return errors.New("invalid items quantity")
	}
	mergeItems := mergItemsQuantity(in.Items)

	// validate stock service
	log.Println(mergeItems)
	return nil
}

func mergItemsQuantity(itemsWithQuantity []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity{
	merged := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range itemsWithQuantity {
		found := false
		for _, finalItem := range merged {
			if finalItem.Id==item.Id {
				finalItem.Quantity+=item.Quantity
				found=true
				break
			}
		}

		if !found {
			merged=append(merged, item)
		}
	}
	return merged
}