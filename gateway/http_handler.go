package main

import "net/http"

type handler struct {
	//gtw
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.handleCreateOrder)
}
func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {

}