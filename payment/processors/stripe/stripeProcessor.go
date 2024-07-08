package stripeProcessor

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

type StripeProcessor struct {

}

func NewStripeProcessor() *StripeProcessor {
	return &StripeProcessor{}
}

func (s *StripeProcessor)CreaterPaymentLink(in *pb.CreateOrderResponse)(string, error) {
	log.Printf("Creationg payment link for order %v\n", in)
	
	var items []*stripe.CheckoutSessionLineItemParams
	gtwSuccessUrl := fmt.Sprintf("%s/success.html?customerId=%s&orderId=%s", gtwAddr, in.Order.CustomerId, in.Order.Id)
	gtwCancelUrl := fmt.Sprintf("%s/cancel.html", gtwAddr)

	for _, item := range in.Order.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			// Price: stripe.String(item.PriceId),
			Price: stripe.String("price_1PYSN0EJEwxXWrvp5iF9aTfD"),
			Quantity: stripe.Int64(item.Quantity),
		})
	}
	params := &stripe.CheckoutSessionParams{
		LineItems: items,
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gtwSuccessUrl),
		CancelURL: stripe.String(gtwCancelUrl),
	  }
	result, err := session.New(params)
	if err != nil {
		log.Printf("link error %v\n", err)
		return "", err	
	}
	return result.URL, nil
}