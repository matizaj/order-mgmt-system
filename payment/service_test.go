package main

import (
	"context"
	"testing"

	"github.com/matizaj/oms/common/discovery/inmemory"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/payment/gateway"
	testprocessor "github.com/matizaj/oms/payment/processors/test_processor"
)

func TestSetvice(t *testing.T) {
	tprocessor := testprocessor.NewTestProcessor()
	inmem_registry := inmemory.NewRegistry()
	gtw := gateway.NewPaymentGateway(inmem_registry)
	service := NewPaymentService(tprocessor, gtw)
	
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