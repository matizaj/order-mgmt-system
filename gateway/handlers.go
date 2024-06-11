package main

import "net/http"

type Handler struct {
	//gateway
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)
}

func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}