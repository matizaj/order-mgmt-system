package main

import "context"

type store struct {
}

func NewOrderStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}