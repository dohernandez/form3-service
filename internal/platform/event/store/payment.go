package store

import (
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/hellofresh/goengine/aggregate"
)

// PaymentType is the type used to identify the payment transaction within the event store
const PaymentType = "Payment"

// NewTransactionPaymentType۰v0 create an aggregate.Type to allow the repository to reconstitute the transaction.Payment
func NewTransactionPaymentType۰v0() (*aggregate.Type, error) {
	return aggregate.NewType(PaymentType, func() aggregate.Root {
		return PaymentState()
	})
}

// PaymentState returns the instance use to marshall and unmarshal state of the payment
func PaymentState() *transaction.Payment {
	return &transaction.Payment{}
}
