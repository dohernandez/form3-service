package payment

import (
	"context"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/event/store"
)

// Response represents the transaction payment response
type Response struct {
	*transaction.Payment
	Type string
}

// encodeToResponse encode transaction to response
func encodeToResponse(_ context.Context, t *transaction.Payment) interface{} {
	var r Response

	r.Payment = t
	r.Type = store.PaymentType

	return &r
}
