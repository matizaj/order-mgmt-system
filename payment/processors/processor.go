package processors

import (
	pb "github.com/matizaj/oms/common/proto"
)
type PaymentProcessor interface {
	CreaterPaymentLink(*pb.Order)(string, error)
}