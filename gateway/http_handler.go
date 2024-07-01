package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	// gtw
	grpcClient pb.OrderServiceClient
}

func NewHttpHandler(grpcClient pb.OrderServiceClient) *handler {
	return &handler{grpcClient}
}

func (h *handler) registerRoutes(router *http.ServeMux) {	
	router.HandleFunc("POST /api/customers/{customerId}/orders", h.getOrdersByCustomerId)
}


func (h *handler) getOrdersByCustomerId(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")
	items := []*pb.ItemsWithQuantity{}
	if err := common.ReadJson(r, &items); err != nil {
		common.ErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if err := validateItems(items); err != nil {
		log.Printf("failed to create order %v\n", err)
		common.ErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}
	
	order :=  &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items: items,
	}
	resp, err := h.grpcClient.CreateOrder(context.Background(),order)
	if err != nil {
		log.Printf("failed to create order %v\n", err)
		errGrpc := status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed to create order %v", err),
		)
		common.ErrorJson(w, http.StatusInternalServerError, errGrpc.Error())
		return
	}
	common.WriteJson(w, http.StatusCreated, resp)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items)<0 {
		return errors.New("invalid items count")
	}

	for _, i := range items {
		if i.Id == "" {
			return errors.New("item Id is required")
		}
		if i.Quantity <=0  {
			return errors.New("invalid item quantity")
		}
	}
	return nil
}