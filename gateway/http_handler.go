package main

import (
	"errors"
	"log"

	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/gateway/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	// gt
	gateway gateway.OrderGateway
}

func NewHttpHandler(gateway gateway.OrderGateway) *handler {
	return &handler{gateway}
}

func (h *handler) registerRoutes(router *http.ServeMux) {	
	router.HandleFunc("POST /api/customers/{customerId}/orders", h.createOrder)
	router.HandleFunc("GET /api/customers/{customerId}/orders/{orderId}", h.getOrdersByCustomerId)
}


func (h *handler) createOrder(w http.ResponseWriter, r *http.Request) {
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

func (h *handler) getOrdersByCustomerId(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")
	orderId := r.PathValue("orderId")

	resp, err := h.gateway.GetOrder(r.Context(), customerId, orderId)
	if err != nil {
		grpcErr := status.Errorf(codes.InvalidArgument,"")
		common.ErrorJson(w, http.StatusBadRequest, grpcErr.Error())
		return
	}
	log.Printf("order %v\n", resp)

	
	
	common.WriteJson(w, http.StatusOK, resp)
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