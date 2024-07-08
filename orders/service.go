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

func (s *service) CreateOrder(ctx context.Context, items []*pb.Item, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
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
	order, err := s.store.Get(ctx, in.CustomerId, in.OrderId)
	if err != nil {
		return nil, err
	}
	o := &pb.GetOrderResponse{
		Order: &pb.Order{
			Id: order.Order.Id,
			CustomerId: order.Order.CustomerId,
			Status: "success",
			Items: order.Order.Items,
			},
		}
	

	return o, nil
}

func (s *service) ValidateOrder(ctx context.Context, in *pb.CreateOrderRequest) ([]*pb.Item, error)	 {
	if len(in.Items) <= 0 {
		return nil, errors.New("invalid items quantity")
	}
	mergeItems := mergItemsQuantity(in.Items)

	// validate stock service
	log.Println(mergeItems)
	items, err := checkStockAvailability(ctx, mergeItems)
	if err != nil {
		return nil, err
	}
	return items, nil
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

func checkStockAvailability(ctx context.Context, mergedItems []*pb.ItemsWithQuantity) ([]*pb.Item, error) {
	items := []*pb.Item{}
	for _, i := range mergedItems {
		items = append(items, &pb.Item{
			Id: i.Id,
			Quantity: i.Quantity,
			PriceId: "test",
			Name: "default name",
		})
	}
	return items, nil
}