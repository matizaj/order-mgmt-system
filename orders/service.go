package main

import "context"

type service struct {
	store OrderStore
}

func NewOrderService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}