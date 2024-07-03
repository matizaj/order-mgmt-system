package main

import (
	"errors"
	"log"

	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/gateway/gateway"
)

type handler struct {
	// gt
	gateway gateway.OrderGateway
}

func NewHttpHandler(gateway gateway.OrderGateway) *handler {
	return &handler{gateway}
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
		log.Printf("failed to validate order %v\n", err)
		common.ErrorJson(w, http.StatusBadRequest, err.Error())
		return
	}
	orderReq := &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items: items,
	}
	resp, err := h.gateway.CreateOrder(r.Context(), orderReq)
	if err != nil {
		log.Printf("failed to validate order %v\n", err)
		common.ErrorJson(w, http.StatusInternalServerError, err.Error())
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