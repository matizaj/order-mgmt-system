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
	order := &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id: "1",
			CustomerId: "43",
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

