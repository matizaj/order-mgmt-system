package main

import (
	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/api"
)

type Handler struct {
	GrpcClient pb.OrderServiceClient
}

func NewHandler(grpcClient pb.OrderServiceClient) *Handler {
	return &Handler{GrpcClient: grpcClient}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.GrpcClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerId,
		Items: items,
	})
	w.Write([]byte("hello"))
}