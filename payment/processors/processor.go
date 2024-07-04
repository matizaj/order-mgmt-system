package processors

import (
	pb "github.com/matizaj/oms/common/proto"
)
type PaymentProcessor interface {
	CreaterPaymentLink(*pb.CreateOrderResponse)(string, error)
}