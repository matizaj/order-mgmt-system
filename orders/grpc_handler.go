package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/matizaj/oms/common/broker"
	pb "github.com/matizaj/oms/common/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type grpcHandler struct {
	pb.OrderServiceServer
	service OrderService
	queue *amqp.Channel
}
func NewGrpcHandler(service OrderService, queue *amqp.Channel) *grpcHandler {
	return &grpcHandler{service: service, queue: queue}
}

func (h *grpcHandler) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("New order received %v\n", in)	
	
	order , err := h.service.CreateOrder(ctx, in)
	if err != nil {
		return nil, err
	}

	q, err := h.queue.QueueDeclare(broker.OrderCreatedEvent, true, false, false,false, nil)
	if err != nil {
		return nil, err
	}

	marshalledOrder, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	log.Printf("order %v, %v", order, marshalledOrder)
	h.queue.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: marshalledOrder,
		DeliveryMode: amqp.Persistent,
	})
	return order, nil
}

func (h *grpcHandler) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	log.Printf("invoking get order with %v", in)
	
	order, err := h.service.GetOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	
	return order, nil

}