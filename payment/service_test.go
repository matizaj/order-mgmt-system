package main

import (
	"context"
	"testing"

	pb "github.com/matizaj/oms/common/proto"
	testprocessor "github.com/matizaj/oms/payment/processors/test_processor"
)

func TestSetvice(t *testing.T) {
	tprocessor := testprocessor.NewTestProcessor()
	service := NewPaymentService(tprocessor)
	
	t.Run("should create a paymewnt link", func(t *testing.T) {
		dymmyOrder := &pb.CreateOrderResponse{}

		link, err := service.CreatePayment(context.Background(), dymmyOrder)
		if err != nil {
			t.Errorf("CreatePayment() error = %v, want nil ", err)
		}
		if link == "" {
			t.Error("Link is empty")
		}
	})
}