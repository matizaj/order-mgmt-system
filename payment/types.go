package main


import (
	"context"
	pb "github.com/matizaj/oms/common/proto"
)

type PaymentService interface {
	CreatePayment(context.Context, *pb.CreateOrderResponse)(string, error)
}