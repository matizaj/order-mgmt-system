package main

import "context"

func main() {
	orderStore := NewStore()
	orderSvc := NewService(orderStore)
	orderSvc.CreateOrder(context.TODO())
}