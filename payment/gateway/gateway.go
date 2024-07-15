package gateway

import (
	"context"

)

type OrderGateway interface {
	UpdateOrderAfterPaymentLink(ctx context.Context, orderId, paymentLink string) error
}