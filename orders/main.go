package main

import "context"

func main() {
	orderStore := NewOrderStore()
	orderSvc := NewOrderService(orderStore)
	orderSvc.CreateOrder(context.Background())
}