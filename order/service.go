package main

import (
	"context"
	"log"
	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, order *pb.CreateOrderRequest) error {
	if len(order.Items)<=0 {
		return common.ErrNoItems
	}
	mergedItems := mergedItemsQuantity(order.Items)
	log.Println("Merged items: ", mergedItems)
	return nil
}

func mergedItemsQuantity(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity{
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item:=range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.Id == item.Id {
				finalItem.Quantity+=item.Quantity
				found=true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
		
	}
	return merged
}