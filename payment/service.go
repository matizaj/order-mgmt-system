package main 

import (
	"context"
	pb "github.com/matizaj/oms/common/proto"
	stripeProcessor "github.com/matizaj/oms/payment/processors/stripe"
)

type service struct {
	// gateway
	stripeProcessor *stripeProcessor.StripeProcessor
}

func NewPaymentService(stripeProcessor *stripeProcessor.StripeProcessor) *service {
	return &service{stripeProcessor}
}

func (s service) CreatePayment(ctx context.Context, in *pb.Order)(string, error) {
	// connect to payment processor
	link, err:=  stripeProcessor.NewStripeProcessor().CreaterPaymentLink(in)
	if err != nil {
		return "", err
	}
	return link, nil
}