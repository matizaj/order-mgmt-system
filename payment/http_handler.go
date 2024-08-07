package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/matizaj/oms/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/webhook"
  pb "github.com/matizaj/oms/common/proto"
)

type PaymentHTTPHandler struct {
	channel *amqp.Channel
}

func NewPaymentHandler(channel *amqp.Channel) *PaymentHTTPHandler {
	return &PaymentHTTPHandler{channel}
}


func (h *PaymentHTTPHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("/webhook", h.handleWebhookCheckout)
  router.HandleFunc("GET /ok", func (w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("ok"))
  })
}

func (h *PaymentHTTPHandler) handleWebhookCheckout(w http.ResponseWriter, r *http.Request) {
  log.Println("webhook")
	const MaxBodyBytes = int64(65536)
  r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

  body, err := io.ReadAll(r.Body)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
    w.WriteHeader(http.StatusServiceUnavailable)
    return
  }

  fmt.Fprintf(os.Stdout, "Got body: %s\n", body)

  
  event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), endpointStripeSecret)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
    w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
    return
  }

  if event.Type == "checkout.session.completed" || event.Type == "charge.succeded" {
    var session stripe.CheckoutSession
    err := json.Unmarshal(event.Data.Raw, &session)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
      w.WriteHeader(http.StatusBadRequest)
      return
    }
  
	if session.PaymentStatus == "paid" {
		log.Printf("Payment for checkout Session %v succeeded", session.ID)

    orderId := session.Metadata["orderId"]
    customerId := session.Metadata["customerId"]
		//publish mesage
    ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
    o := &pb.Order{
      Id: orderId,
      CustomerId: customerId,
      Status: "paid",
      PaymentLink: "",
    }
    marshalledOrder, _ := json.Marshal(o)
    defer cancel()
    h.channel.PublishWithContext(ctx, broker.OrderPaidEvent, "", false, false, amqp.Publishing{
      ContentType: "application/json",
      Body: marshalledOrder,
      DeliveryMode: amqp.Persistent,
    })
	}
}
  w.WriteHeader(http.StatusOK)

}

