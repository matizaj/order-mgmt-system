package main

import "context"

type store struct {
	// mongo
}

func NewStore() *store {
	return &store{}
}

func (s *store) CreateStore(context.Context) error {
	return nil
}