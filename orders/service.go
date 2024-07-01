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

func (s *service) CreateOrder(ctx context.Context) error {
	return nil
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