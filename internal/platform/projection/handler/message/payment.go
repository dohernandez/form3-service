package message

import (
	"context"
	"errors"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/pkg/log"
	"github.com/hellofresh/goengine"
	"github.com/hellofresh/goengine/aggregate"
)

// PaymentCreator defines the way to create a payment
type PaymentCreator interface {
	Create(
		ctx context.Context,
		ID aggregate.ID,
		version transaction.Version,
		organisationID transaction.OrganisationID,
		attributes interface{},
	) error
}

// TransactionPaymentCreatedHandler۰v0 creates a payment handler for notifier messages
//
// This handler handle all those state messages trigger when a new payment is create.
// The function is used to update the payment projection
func TransactionPaymentCreatedHandler۰v0(creator PaymentCreator) goengine.MessageHandler {
	return func(ctx context.Context, state interface{}, message goengine.Message) (interface{}, error) {
		writeDebugMessage(ctx, "transaction_payment_created_v0")

		event, ok := message.Payload().(transaction.PaymentCreated۰v0)
		if !ok {
			return nil, errors.New("wrong message version")
		}

		if err := creator.Create(
			ctx,
			event.ID,
			transaction.Version0,
			event.OrganisationID,
			event.Attributes,
		); err != nil {
			return nil, err
		}

		return state, nil
	}
}

func writeDebugMessage(ctx context.Context, event string) {
	logger := log.FromContext(ctx)
	if logger != nil {
		logger.Debugf("handling event %s", event)
	}
}
