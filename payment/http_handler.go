package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

func NewPaymentHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel}
}

func (h *PaymentHTTPHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/weebhook", h.handleWebhookCheckout)
}

func (h *PaymentHTTPHandler) handleWebhookCheckout(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
  req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)

  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
    w.WriteHeader(http.StatusServiceUnavailable)
    return
  }

  fmt.Fprintf(os.Stdout, "Got body: %s\n", body)

  w.WriteHeader(http.StatusOK)
}