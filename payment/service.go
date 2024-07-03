package main 

import (
	"context"
	pb "github.com/matizaj/oms/common/proto"
)

type service struct {
	// gateway
}

func NewPaymentService() *service {
	return &service{}
}

func (s service) CreatePayment(context.Context, *pb.Order)(string, error) {
	// connect to payment processor
	return "", nil
}