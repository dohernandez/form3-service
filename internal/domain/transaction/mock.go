package transaction

import (
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/google/uuid"
	"github.com/hellofresh/goengine/aggregate"
)

// NewPersonMock creates a Person mock for tests
// nolint:unused
func NewPersonMock() Person {
	return Person{
		Name:    "Emelia Jane Brown",
		Address: "10 Debtor Crescent Sourcetown NE1",
	}
}

// NewBankAccountMock creates a BankAccount mock for tests
// nolint:unused
func NewBankAccountMock() BankAccount {
	person := NewPersonMock()
	account := NewAccountMock()

	return BankAccount{
		Person:            person,
		Account:           account,
		AccountName:       "EJ Brown Black",
		AccountNumberCode: "IBAN",
	}
}

// NewAccountMock creates a Account mock for tests
// nolint:unused
func NewAccountMock() Account {
	return Account{
		Bank: Bank{
			ID:     203301,
			IDCode: "GBDSC",
		},
		Number: "GB29XABC10161234567801",
		Type:   0,
	}
}

// NewChargesInformationMock creates a ChargesInformation mock for tests
// nolint:unused
func NewChargesInformationMock() ChargesInformation {
	return ChargesInformation{
		BearerCode: "SHAR",
		SenderCharges: []SenderCharge{
			{
				Amount:   5,
				Currency: "GBP",
			},
			{
				Amount:   10,
				Currency: "USD",
			},
		},
		ReceiverChargesAmount:   10,
		ReceiverChargesCurrency: "USD",
	}
}

// NewSenderChargeMock creates a SenderCharge mock for tests
// nolint:unused
func NewSenderChargeMock() SenderCharge {
	return SenderCharge{
		Amount:   5,
		Currency: "GBP",
	}
}

// NewPayment۰v0Mock creates a Payment۰v0 mock for tests
// nolint:unused
func NewPayment۰v0Mock() Payment۰v0 {
	beneficiaryParty := NewBankAccountMock()
	chargesInformation := NewChargesInformationMock()
	debtorParty := NewBankAccountMock()
	fx := NewFxMock()
	processingDate := NewProcessingDateMock()
	sponsorParty := NewAccountMock()

	return Payment۰v0{
		Amount:               100.21,
		BeneficiaryParty:     beneficiaryParty,
		ChargesInformation:   chargesInformation,
		Currency:             "GBP",
		DebtorParty:          debtorParty,
		EndToEndReference:    "Wil piano Jan",
		FX:                   fx,
		NumericReference:     1002001,
		PaymentID:            123456789012345678,
		PaymentPurpose:       "Paying for goods/services",
		PaymentScheme:        "FPS",
		PaymentType:          "Credit",
		ProcessingDate:       processingDate,
		Reference:            "Payment for Em's piano lessons",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",
		SponsorParty:         sponsorParty,
	}
}

// NewFxMock creates a FX mock for tests
// nolint:unused
func NewFxMock() FX {
	return FX{
		ContractReference: "FX123",
		ExchangeRate:      2.0,
		Amount:            200.42,
		Currency:          "USD",
	}
}

// NewProcessingDateMock creates a ProcessingDate mock for tests
// nolint:unused
func NewProcessingDateMock() ProcessingDate {
	return ProcessingDate{
		must.ParseTime("2017-01-18", ProcessingDateLayout),
	}
}

// NewPaymentCreated۰v0Mock creates a PaymentCreated۰v0 mock for tests
// nolint:unused
func NewPaymentCreated۰v0Mock() PaymentCreated۰v0 {
	aggregateID := aggregate.GenerateID()
	organisationID := OrganisationID(uuid.New().String())
	payment۰v0 := NewPayment۰v0Mock()

	return PaymentCreated۰v0{
		ID:             aggregateID,
		OrganisationID: organisationID,
		Attributes:     payment۰v0,
	}
}

// NewPaymentBeneficiaryUpdated۰v0Mock creates a PaymentBeneficiaryUpdated۰v0 mock for tests
// nolint:unused
func NewPaymentBeneficiaryUpdated۰v0Mock() PaymentBeneficiaryUpdated۰v0 {
	aggregateID := aggregate.GenerateID()
	beneficiary := NewBankAccountMock()

	return PaymentBeneficiaryUpdated۰v0{
		ID:          aggregateID,
		Beneficiary: beneficiary,
	}
}

// NewPaymentDeletedMock creates a PaymentDeleted mock for tests
// nolint:unused
func NewPaymentDeletedMock() PaymentDeleted {
	aggregateID := aggregate.GenerateID()

	return PaymentDeleted{
		ID: aggregateID,
	}
}
