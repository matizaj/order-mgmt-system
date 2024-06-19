package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/api"
	"github.com/matizaj/oms/gateway/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gtw gateway.OrdersGateway
}

func NewHandler(gtw gateway.OrdersGateway) *Handler {
	return &Handler{gtw}
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
	if err := validateItems(items); err != nil {
		log.Println("the error", err)
		common.WriteError(w, http.StatusBadRequest, "payload validation failed")
		return
	}

	order, err := h.gtw.CreateOrder(r.Context(),&pb.CreateOrderRequest{
		CustomerId: customerId,
		Items: items,
	})
	errStatus := status.Convert(err)
	if errStatus != nil {
		log.Print("some grpc erros")
		if errStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, errStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, "failed to create order")
		return
	}

	common.WriteJson(w, http.StatusCreated, order)
}

func validateItems(items []*pb.ItemsWithQuantity) error{
	log.Println("gateway validator payload")
	if len(items)==0 {
		return common.ErrNoItems
	}

	for _, i :=range items {
		if i.Id == "" {
			return errors.New("invalid item")
		}
		if i.Quantity <= 0 {
			return errors.New("invalid quantity for item")
		}
		
	}
	return nil
}