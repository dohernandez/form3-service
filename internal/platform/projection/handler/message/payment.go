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

// PaymentCreatedHandler۰v0 creates a payment handler for notifier messages
//
// This handler handle all those state messages trigger when a new payment is created.
// The function is used to create the payment in the projection
func PaymentCreatedHandler۰v0(creator PaymentCreator) goengine.MessageHandler {
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

// PaymentBeneficiaryUpdater defines the way to update the payment's beneficiary
type PaymentBeneficiaryUpdater interface {
	UpdateBeneficiary(ctx context.Context, ID aggregate.ID, beneficiary transaction.BankAccount) error
}

// PaymentDeleter defines the way to delete the payment's beneficiary
type PaymentDeleter interface {
	Delete(ctx context.Context, ID aggregate.ID) error
}

// PaymentBeneficiaryUpdatedHandler۰v0 updates a payment's beneficiary handler for notifier messages
//
// This handler handle all those state messages trigger when a payment's beneficiary is updated.
// The function is used to update the payment in the projection
func PaymentBeneficiaryUpdatedHandler۰v0(updater PaymentBeneficiaryUpdater) goengine.MessageHandler {
	return func(ctx context.Context, state interface{}, message goengine.Message) (interface{}, error) {
		writeDebugMessage(ctx, "transaction_payment_beneficiary_updated_v0")

		event, ok := message.Payload().(transaction.PaymentBeneficiaryUpdated۰v0)
		if !ok {
			return nil, errors.New("wrong message version")
		}

		if err := updater.UpdateBeneficiary(
			ctx,
			event.ID,
			event.Beneficiary,
		); err != nil {
			return nil, err
		}

		return state, nil
	}
}

// PaymentDeletedHandler updates a payment's beneficiary handler for notifier messages
//
// This handler handle all those state messages trigger when a payment is deleted.
// The function is used to delete the payment in the projection
func PaymentDeletedHandler(deleter PaymentDeleter) goengine.MessageHandler {
	return func(ctx context.Context, state interface{}, message goengine.Message) (interface{}, error) {
		writeDebugMessage(ctx, "transaction_payment_deleted")

		event, ok := message.Payload().(transaction.PaymentDeleted)
		if !ok {
			return nil, errors.New("wrong message version")
		}

		if err := deleter.Delete(
			ctx,
			event.ID,
		); err != nil {
			return nil, err
		}

		return state, nil
	}
}
