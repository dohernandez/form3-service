package message_test

import (
	"context"
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/projection/handler/message"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTransactionPaymentCreatedHandler۰v0(t *testing.T) {
	paymentCreated۰v0 := transaction.NewPaymentCreated۰v0Mock()

	testCases := []struct {
		scenario string

		createFunc func(
			ID aggregate.ID,
			version transaction.Version,
			organisationID transaction.OrganisationID,
			attributes interface{},
		) error
		payloadFunc func() interface{}

		err error
	}{
		{
			scenario: "Create payment in the projection successful",
			createFunc: func(
				ID aggregate.ID,
				version transaction.Version,
				organisationID transaction.OrganisationID,
				attributes interface{},
			) error {
				assert.Equal(t, paymentCreated۰v0.ID, ID)
				assert.Equal(t, transaction.Version0, version)
				assert.Equal(t, paymentCreated۰v0.OrganisationID, organisationID)
				assert.Equal(t, paymentCreated۰v0.Attributes, attributes)

				return nil
			},
			payloadFunc: func() interface{} {
				return paymentCreated۰v0
			},
		},
		{
			scenario: "Create payment in the projection unsuccessful, wrong payload",
			createFunc: func(
				ID aggregate.ID,
				version transaction.Version,
				organisationID transaction.OrganisationID,
				attributes interface{},
			) error {
				panic("should not be called")
			},
			payloadFunc: func() interface{} {
				return "wrong payload"
			},
			err: errors.New("wrong message version"),
		},
		{
			scenario: "Create payment in the projection unsuccessful, creator error",
			createFunc: func(
				ID aggregate.ID,
				version transaction.Version,
				organisationID transaction.OrganisationID,
				attributes interface{},
			) error {

				assert.Equal(t, paymentCreated۰v0.ID, ID)
				assert.Equal(t, transaction.Version0, version)
				assert.Equal(t, paymentCreated۰v0.OrganisationID, organisationID)
				assert.Equal(t, paymentCreated۰v0.Attributes, attributes)

				return errors.New("creator error")
			},
			payloadFunc: func() interface{} {
				return paymentCreated۰v0
			},
			err: errors.New("creator error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			creator := message.NewCallbackPaymentCreatorMock(tc.createFunc)
			handler := message.TransactionPaymentCreatedHandler۰v0(creator)

			message := message.NewMessageMock(tc.payloadFunc)

			_, err := handler(context.TODO(), true, message)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTransactionPaymentBeneficiaryUpdatedHandler۰v0(t *testing.T) {
	paymentBeneficiaryUpdated۰v0 := transaction.NewPaymentBeneficiaryUpdated۰v0Mock()

	testCases := []struct {
		scenario string

		beneficiaryUpdaterFunc func(ID aggregate.ID, beneficiary transaction.BankAccount) error
		payloadFunc            func() interface{}

		err error
	}{
		{
			scenario: "Update beneficiary in the payment projection successful",
			beneficiaryUpdaterFunc: func(ID aggregate.ID, beneficiary transaction.BankAccount) error {
				assert.Equal(t, paymentBeneficiaryUpdated۰v0.ID, ID)
				assert.Equal(t, paymentBeneficiaryUpdated۰v0.Beneficiary, beneficiary)

				return nil
			},
			payloadFunc: func() interface{} {
				return paymentBeneficiaryUpdated۰v0
			},
		},
		{
			scenario: "Update beneficiary in the payment projection unsuccessful, wrong payload",
			beneficiaryUpdaterFunc: func(ID aggregate.ID, beneficiary transaction.BankAccount) error {
				panic("should not be called")
			},
			payloadFunc: func() interface{} {
				return "wrong payload"
			},
			err: errors.New("wrong message version"),
		},
		{
			scenario: "Update beneficiary in the payment projection unsuccessful, updater error",
			beneficiaryUpdaterFunc: func(ID aggregate.ID, beneficiary transaction.BankAccount) error {
				assert.Equal(t, paymentBeneficiaryUpdated۰v0.ID, ID)
				assert.Equal(t, paymentBeneficiaryUpdated۰v0.Beneficiary, beneficiary)

				return errors.New("creator error")
			},
			payloadFunc: func() interface{} {
				return paymentBeneficiaryUpdated۰v0
			},
			err: errors.New("creator error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			updater := message.NewCallbackPaymentBeneficiaryUpdaterMock(tc.beneficiaryUpdaterFunc)
			handler := message.TransactionPaymentBeneficiaryUpdatedHandler۰v0(updater)

			message := message.NewMessageMock(tc.payloadFunc)

			_, err := handler(context.TODO(), true, message)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
