package stripe

import (
	"fmt"
	"log"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

var (
	gtwAddr = common.EnvString("GTW_ADDR", "")
)

type stripeProcessor struct {

}

func NewStripeProcessor() *stripeProcessor {
	return &stripeProcessor{}
}

func (s *stripeProcessor)CreaterPaymentLink(in *pb.Order)(string, error) {
	log.Printf("Creationg payment link for order %v\n", in)
	log.Printf("gtw address %v\n", gtwAddr)
	domain:= "https://example.com"
	
	var items []*stripe.CheckoutSessionLineItemParams
	gatewaySuccessUrl:=fmt.Sprintf("%s/success.html", gtwAddr)

	for _, item := range in.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price: stripe.String(item.PriceId),
			Quantity: stripe.Int64(item.Quantity),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:items,
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessUrl),
		CancelURL: stripe.String(domain + "/cancel.html"),
	  }
	result, err := session.New(params)
	if err != nil {
		return "", nil	
	}
	return result.URL, nil
}