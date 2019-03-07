package transaction_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestChargesInformationValidate(t *testing.T) {
	testCases := []struct {
		scenario               string
		chargesInformationFunc func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation
		err                    error
	}{
		{
			scenario: "Charges information is valid",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				return chargesInformation
			},
		},
		{
			scenario: "Charges information is invalid due to empty bearer code",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				chargesInformation.BearerCode = ""

				return chargesInformation
			},
			err: errors.Errorf("bearer_code: cannot be blank."),
		},
		{
			scenario: "Charges information is invalid due to empty sender charges",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				chargesInformation.SenderCharges = []transaction.SenderCharge{}

				return chargesInformation
			},
			err: errors.Errorf("sender_charges: cannot be empty."),
		},
		{
			scenario: "Charges information is invalid due to empty currency in the first sender charges",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				chargesInformation.SenderCharges[0].Currency = ""

				return chargesInformation
			},
			err: errors.Errorf("sender_charges: (0: (currency: cannot be blank.).)."),
		},
		{
			scenario: "Charges information is invalid due to zero receiver charges amount",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				chargesInformation.ReceiverChargesAmount = 0

				return chargesInformation
			},
			err: errors.Errorf("receiver_charges_amount: cannot be zero."),
		},
		{
			scenario: "Charges information is invalid due to empty receiver charges currency",
			chargesInformationFunc: func(chargesInformation transaction.ChargesInformation) transaction.ChargesInformation {
				chargesInformation.ReceiverChargesCurrency = ""

				return chargesInformation
			},
			err: errors.Errorf("receiver_charges_currency: cannot be blank."),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			chargesInformation := tc.chargesInformationFunc(transaction.NewChargesInformationMock())

			err := chargesInformation.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChargeValidate(t *testing.T) {
	testCases := []struct {
		scenario         string
		senderChargeFunc func(senderCharge transaction.SenderCharge) transaction.SenderCharge
		err              error
	}{
		{
			scenario: "Sender charge is valid",
			senderChargeFunc: func(senderCharge transaction.SenderCharge) transaction.SenderCharge {
				return senderCharge
			},
		},
		{
			scenario: "Sender charge is invalid due to empty currency",
			senderChargeFunc: func(senderCharge transaction.SenderCharge) transaction.SenderCharge {
				senderCharge.Currency = ""

				return senderCharge
			},
			err: errors.Errorf("currency: cannot be blank."),
		},
		{
			scenario: "Sender charge is invalid due to zero amount",
			senderChargeFunc: func(senderCharge transaction.SenderCharge) transaction.SenderCharge {
				senderCharge.Amount = 0

				return senderCharge
			},
			err: errors.Errorf("amount: cannot be zero."),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			senderCharge := tc.senderChargeFunc(transaction.NewSenderChargeMock())

			err := senderCharge.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
