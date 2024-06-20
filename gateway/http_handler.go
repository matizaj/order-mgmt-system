package main

import (
	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/api"
)

type handler struct {
	orderClient pb.OrderServiceClient
}

func NewHandler(orderClient pb.OrderServiceClient) *handler {
	return &handler{orderClient}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.handleCreateOrder)
}
func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var items []*pb.ItemsWithQuantity

	if err := common.ReadJson(r, &items); err != nil {
		common.ErrorJson(w, http.StatusBadRequest, err.Error())
	}

	h.orderClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: r.PathValue("customerId"),
		Items: items,
	})
}