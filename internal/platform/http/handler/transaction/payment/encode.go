package payment

import (
	"context"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/event/store"
)

// Response represents the transaction payment response
type Response struct {
	*transaction.Payment
	Type string `json:"type"`
}

// encodeToResponse encode transaction to response
func encodeToResponse(_ context.Context, t *transaction.Payment) interface{} {
	var r Response

	r.Payment = t
	r.Type = store.PaymentType

	return &r
}

// AllResponse represents the transaction payment response
type AllResponse struct {
	Data  []Response `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

// encodeToResponse encode transaction to response
func encodeAllToResponse(_ context.Context, ts []*transaction.Payment, url string) interface{} {
	var r AllResponse

	for _, t := range ts {
		var rItem Response

		rItem.Payment = t
		rItem.Type = store.PaymentType

		r.Data = append(r.Data, rItem)
	}

	r.Links.Self = url

	return &r
}
