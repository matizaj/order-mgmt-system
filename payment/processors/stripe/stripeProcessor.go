package stripeProcessor

import (
	"log"

	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

var (
	gtwAddr = common.EnvString("GTW_ADDR", "")
)

type StripeProcessor struct {

}

func NewStripeProcessor() *StripeProcessor {
	return &StripeProcessor{}
}

func (s *StripeProcessor)CreaterPaymentLink(in *pb.CreateOrderResponse)(string, error) {
	log.Printf("Creationg payment link for order %v\n", in)
	log.Printf("gtw address %v\n", gtwAddr)
	
	var items []*stripe.CheckoutSessionLineItemParams

	for _, item := range in.Order.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			// Price: stripe.String(item.PriceId),
			Price: stripe.String("price_1PYSN0EJEwxXWrvp5iF9aTfD"),
			Quantity: stripe.Int64(item.Quantity),
		})
	}
	log.Printf("Line ITEMS %v\n", items)
	params := &stripe.CheckoutSessionParams{
		LineItems: items,
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gtwAddr + "/success.html"),
		CancelURL: stripe.String(gtwAddr + "/cancel.html"),
	  }
	result, err := session.New(params)
	if err != nil {
		log.Printf("l;ink error %v\n", err)
		return "", err	
	}
	return result.URL, nil
}