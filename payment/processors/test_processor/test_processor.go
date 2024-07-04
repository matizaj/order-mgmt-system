package testprocessor

import 	pb "github.com/matizaj/oms/common/proto"

type InMemoryTestProcessor struct {

}

func NewTestProcessor() *InMemoryTestProcessor {
	return &InMemoryTestProcessor{}
}

func (in *InMemoryTestProcessor) CreaterPaymentLink(*pb.CreateOrderResponse) (string, error) {
	return "test_link", nil
}