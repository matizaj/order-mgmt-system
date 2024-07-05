package main

import "context"

type store struct {
}

func NewStore() *store {
	return &store{}
}
var inMemoryOrders []Order

func (s *store) Create(ctx context.Context, order Order) error {
	order.Order.Id = "123456"
	inMemoryOrders = append(inMemoryOrders, order)
	return nil
}

func (s *store) Get(ctx context.Context, customerId, orderId string) *Order {
	for _, o := range inMemoryOrders {
		if o.Order.CustomerId == customerId && o.Order.Id == orderId {
			return &Order{Order: o.Order}
		}
	}
	return nil
}