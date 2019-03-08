package transaction_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPayment۰v0Validate(t *testing.T) {
	testCases := []struct {
		scenario    string
		paymentFunc func(payment transaction.Payment۰v0) transaction.Payment۰v0
		err         error
	}{
		{
			scenario: "Payment۰v0 is valid",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				return payment
			},
		},
		{
			scenario: "Payment۰v0 is invalid due to beneficiary party is not present",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.BeneficiaryParty = transaction.BankAccount{}

				return payment
			},
			err: errors.Errorf(
				"beneficiary_party: (%s %s %s %s %s %s %s).",
				"account_name: cannot be blank;",
				"account_number: cannot be blank;",
				"account_number_code: cannot be blank;",
				"address: cannot be blank;",
				"bank_id: cannot be zero;",
				"bank_id_code: cannot be blank;",
				"name: cannot be blank.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to charges information is not present",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.ChargesInformation = transaction.ChargesInformation{}

				return payment
			},
			err: errors.Errorf(
				"charges_information: (%s %s %s %s).",
				"bearer_code: cannot be blank;",
				"receiver_charges_amount: cannot be zero;",
				"receiver_charges_currency: cannot be blank;",
				"sender_charges: cannot be empty.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to empty currency",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.Currency = ""

				return payment
			},
			err: errors.Errorf("currency: cannot be blank."),
		},
		{
			scenario: "Payment۰v0 is invalid due to debtor party is not present",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.DebtorParty = transaction.BankAccount{}

				return payment
			},
			err: errors.Errorf(
				"debtor_party: (%s %s %s %s %s %s %s).",
				"account_name: cannot be blank;",
				"account_number: cannot be blank;",
				"account_number_code: cannot be blank;",
				"address: cannot be blank;",
				"bank_id: cannot be zero;",
				"bank_id_code: cannot be blank;",
				"name: cannot be blank.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to empty end to end reference",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.EndToEndReference = ""

				return payment
			},
			err: errors.Errorf("end_to_end_reference: cannot be blank."),
		},
		{
			scenario: "Payment۰v0 is invalid due to fx is not present",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.FX = transaction.FX{}

				return payment
			},
			err: errors.Errorf(
				"fx: (%s %s %s %s).",
				"contract_reference: cannot be blank;",
				"exchange_rate: cannot be zero;",
				"original_amount: cannot be zero;",
				"original_currency: cannot be blank.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to zero numeric reference",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.NumericReference = 0

				return payment
			},
			err: errors.Errorf("numeric_reference: cannot be zero."),
		},
		{
			scenario: "Payment۰v0 is invalid due to zero numeric reference",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.Reference = ""

				return payment
			},
			err: errors.Errorf("reference: cannot be blank."),
		},
		{
			scenario: "Payment۰v0 is invalid due to sponsor party is not present",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.SponsorParty = transaction.Account{}

				return payment
			},
			err: errors.Errorf(
				"sponsor_party: (%s %s %s).",
				"account_number: cannot be blank;",
				"bank_id: cannot be zero;",
				"bank_id_code: cannot be blank.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to empty payment id, purpose, scheme and type",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.PaymentID = 0
				payment.PaymentPurpose = ""
				payment.PaymentScheme = ""
				payment.PaymentType = ""

				return payment
			},
			err: errors.Errorf(
				"%s %s %s %s",
				"payment_id: cannot be blank;",
				"payment_purpose: cannot be blank;",
				"payment_scheme: cannot be blank;",
				"payment_type: cannot be blank.",
			),
		},
		{
			scenario: "Payment۰v0 is invalid due to empty scheme payment type and scheme payment sub type",
			paymentFunc: func(payment transaction.Payment۰v0) transaction.Payment۰v0 {
				payment.SchemePaymentType = ""
				payment.SchemePaymentSubType = ""

				return payment
			},
			err: errors.Errorf(
				"%s %s",
				"scheme_payment_sub_type: cannot be blank;",
				"scheme_payment_type: cannot be blank.",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			payment := tc.paymentFunc(transaction.NewPayment۰v0Mock())

			err := payment.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
