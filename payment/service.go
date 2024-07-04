package main

import (
	"context"
	"log"

	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/payment/processors"
	stripeProcessor "github.com/matizaj/oms/payment/processors/stripe"
)

type service struct {
	// gateway
	stripeProcessor processors.PaymentProcessor
}

func NewPaymentService(stripeProcessor processors.PaymentProcessor) *service {
	return &service{stripeProcessor}
}

func (s service) CreatePayment(ctx context.Context, in *pb.CreateOrderResponse)(string, error) {
	// connect to payment processor
	link, err:=  stripeProcessor.NewStripeProcessor().CreaterPaymentLink(in)
	if err != nil {
		log.Printf("stripe failed %v\n", err)
		return "", err
	}
	return link, nil
}