package main

import (
	"context"
	"log"

	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/payment/gateway"
	"github.com/matizaj/oms/payment/processors"
)

type service struct {
	// gateway
	stripeProcessor processors.PaymentProcessor
	gateway gateway.OrderGateway
}

func NewPaymentService(stripeProcessor processors.PaymentProcessor, gateway gateway.OrderGateway) *service {
	return &service{stripeProcessor, gateway}
}

func (s *service) CreatePayment(ctx context.Context, in *pb.CreateOrderResponse)(string, error) {
	// connect to payment processor
	link, err:= s.stripeProcessor.CreaterPaymentLink(in)
	if err != nil {
		log.Printf("stripe failed %v\n", err)
		return "", err
	}
	err = s.gateway.UpdateOrderAfterPaymentLink(ctx, in.Order.Id, link)
	if err != nil {
		return "", err
	}
	return link, nil
}