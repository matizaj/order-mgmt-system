package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/matizaj/oms/common/proto"

	"github.com/matizaj/oms/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	service PaymentService
}

func NewConsumer(service PaymentService) *consumer {
	return &consumer{service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false,false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	var forever chan struct {}
	go func(){
		for msg := range msgs {
			log.Printf("Reveived message %v\n", msg.Body)
			o:= &pb.CreateOrderResponse{}

			if err := json.Unmarshal(msg.Body, o); err != nil {
				log.Printf("failed to unmarshal order %v\n", err)
				continue
			} 
			log.Printf("ORDER %v\n", o)
			paymentLint, err := c.service.CreatePayment(context.Background(), o)
			if err != nil {
				log.Printf("failed to create payment %v\n", err)
				continue
			}

			log.Printf("Payment link %s\n", paymentLint)
		}
	}()
	<-forever
}