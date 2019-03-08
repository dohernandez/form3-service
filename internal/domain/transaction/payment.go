package transaction

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/hellofresh/goengine/aggregate"
)

type (
	// Payment۰v0 represents payment attributes version 0
	Payment۰v0 struct {
		Amount             float64            `json:"amount,string"`
		BeneficiaryParty   BankAccount        `json:"beneficiary_party"`
		ChargesInformation ChargesInformation `json:"charges_information"`
		Currency           Currency           `json:"currency"`
		DebtorParty        BankAccount        `json:"debtor_party"`
		EndToEndReference  EndToEndReference  `json:"end_to_end_reference"`
		FX                 FX                 `json:"fx"`
		NumericReference   int64              `json:"numeric_reference,string"`
		ProcessingDate     ProcessingDate     `json:"processing_date"`
		Reference          Reference          `json:"reference"`
		SponsorParty       Account            `json:"sponsor_party"`

		PaymentID      int64  `json:"payment_id,string"`
		PaymentPurpose string `json:"payment_purpose"`
		PaymentScheme  Scheme `json:"payment_scheme"`
		PaymentType    Type   `json:"payment_type"`

		SchemePaymentType    SchemePaymentType    `json:"scheme_payment_type"`
		SchemePaymentSubType SchemePaymentSubType `json:"scheme_payment_sub_type"`
	}
)

// Validate validates Payment۰v0
func (p Payment۰v0) Validate() error {
	return validation.ValidateStruct(
		&p,
		// Amount cannot be empty
		validation.Field(&p.Amount, validation.Required.Error("cannot be zero")),
		// Validate BeneficiaryParty using BankAccount validation rules
		validation.Field(&p.BeneficiaryParty),
		// Validate ChargesInformation using ChargesInformation validation rules
		validation.Field(&p.ChargesInformation),
		// Validate Currency using Currency validation rules
		validation.Field(&p.Currency),
		// Validate DebtorParty using BankAccount validation rules
		validation.Field(&p.DebtorParty),
		// Validate EndToEndReference using EndToEndReference validation rules
		validation.Field(&p.EndToEndReference),
		// Validate FX using FX validation rules
		validation.Field(&p.FX),
		// NumericReference cannot be empty
		validation.Field(&p.NumericReference, validation.Required.Error("cannot be zero")),
		// Validate ProcessingDate using ProcessingDate validation rules
		validation.Field(&p.ProcessingDate),
		// Validate Reference using Reference validation rules
		validation.Field(&p.Reference),
		// Validate SponsorParty using Account validation rules
		validation.Field(&p.SponsorParty),
		// PaymentID cannot be empty
		validation.Field(&p.PaymentID, validation.Required),
		// PaymentPurpose cannot be empty
		validation.Field(&p.PaymentPurpose, validation.Required),
		// Validate PaymentScheme using Scheme validation rules
		validation.Field(&p.PaymentScheme),
		// Validate PaymentType using Type validation rules
		validation.Field(&p.PaymentType, validation.Required),
		// Validate SchemePaymentType using Type validation rules
		validation.Field(&p.SchemePaymentType, validation.Required),
		// Validate SchemePaymentSubType using Type validation rules
		validation.Field(&p.SchemePaymentSubType, validation.Required),
	)
}

// Scan implements the Scanner interface.
func (p *Payment۰v0) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var source []byte
	switch v := src.(type) {
	case string:
		source = []byte(v)
	case []byte:
		source = v
	default:
		return errors.New("incompatible type for transaction.Payment۰v0")
	}

	return json.Unmarshal(source, p)

}

// Value implements the driver Valuer interface.
func (p Payment۰v0) Value() (driver.Value, error) {
	if reflect.DeepEqual(p, Payment۰v0{}) {
		return driver.Value([]byte("{}")), nil
	}

	v, err := json.Marshal(&p)

	return driver.Value(v), err
}

// ProcessingDateLayout is the format date to parser ProcessingDate
const ProcessingDateLayout = "2006-01-02"

// ProcessingDate represents the date the payment was processed in the format 2006-01-02
type ProcessingDate struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// Error result is always nil to tolerate faulty input
// nolint:unparam
func (d ProcessingDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(ProcessingDateLayout))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *ProcessingDate) UnmarshalJSON(data []byte) error {
	var processingDate string

	if err := json.Unmarshal(data, &processingDate); err != nil {
		return err
	}

	v, err := time.Parse(ProcessingDateLayout, processingDate)
	if err != nil {
		return err
	}

	d.Time = v

	return nil
}

// Validate validates ProcessingDate cannot be empty
func (d ProcessingDate) Validate() error {
	return validation.Validate(&d.Time, validation.Required)
}

// EndToEndReference represents end to end reference
type EndToEndReference string

// Validate validates EndToEndReference cannot be empty
func (r EndToEndReference) Validate() error {
	v := string(r)

	return validation.Validate(&v, validation.Required)
}

// FX represents fix
type FX struct {
	ContractReference string   `json:"contract_reference"`
	ExchangeRate      float64  `json:"exchange_rate,string"`
	Amount            float64  `json:"original_amount,string"`
	Currency          Currency `json:"original_currency"`
}

// Validate validates FX
func (fx FX) Validate() error {
	return validation.ValidateStruct(
		&fx,
		// ContractReference cannot be empty
		validation.Field(&fx.ContractReference, validation.Required),
		// ExchangeRate cannot be empty
		validation.Field(&fx.ExchangeRate, validation.Required.Error("cannot be zero")),
		// Amount cannot be empty
		validation.Field(&fx.Amount, validation.Required.Error("cannot be zero")),
		// Validate Currency using Currency validation rules
		validation.Field(&fx.Currency),
	)
}

// Reference represents reference
type Reference string

// Validate validates Reference cannot be empty
func (r Reference) Validate() error {
	v := string(r)

	return validation.Validate(&v, validation.Required)
}

// Scheme represents payment scheme
type Scheme string

// Validate validates Scheme cannot be empty
func (s Scheme) Validate() error {
	v := string(s)

	return validation.Validate(&v, validation.Required)
}

// Type represents payment type
type Type string

// Validate validates payment type cannot be empty
func (t Type) Validate() error {
	v := string(t)

	return validation.Validate(&v, validation.Required)
}

// SchemePaymentType represents scheme payment Type
type SchemePaymentType string

// Validate validates SchemePaymentType cannot be empty
func (spt SchemePaymentType) Validate() error {
	v := string(spt)

	return validation.Validate(&v, validation.Required)
}

// SchemePaymentSubType represents scheme payment sub Type
type SchemePaymentSubType string

// Validate validates SchemePaymentSubType cannot be empty
func (spst SchemePaymentSubType) Validate() error {
	v := string(spst)

	return validation.Validate(&v, validation.Required)
}

type (
	// Payment represents payment aggregate root
	Payment struct {
		aggregate.BaseRoot
		ID aggregate.ID `json:"id" db:"id"`

		Version        Version        `json:"version" db:"version"`
		OrganisationID OrganisationID `json:"organisation_id" db:"organisation_id"`
		Attributes     interface{}    `json:"attributes" db:"attributes"`
	}

	// OrganisationID represents an UUID
	OrganisationID string
	// Version represents the version of the payment structure
	Version uint
)

// Version0 version 0 of the payment structure
const Version0 = Version(0)

// AggregateID returns the payment aggregate.ID
func (p *Payment) AggregateID() aggregate.ID {
	return p.ID
}

// Apply changes to the Payment
func (p *Payment) Apply(event *aggregate.Changed) {
	if event, ok := event.Payload().(PaymentCreated۰v0); ok {
		p.applyPaymentCreated۰v0(event)
	}
}

// applyPaymentCreated۰v0 applies the creation of the payment in its version 0
func (p *Payment) applyPaymentCreated۰v0(paymentCreated PaymentCreated۰v0) {
	p.ID = paymentCreated.ID
	p.Version = Version0
	p.OrganisationID = paymentCreated.OrganisationID
	p.Attributes = paymentCreated.Attributes
}

// PaymentCreated۰v0 a DomainEvent indicating that payment was created in its version 0
type PaymentCreated۰v0 struct {
	ID             aggregate.ID   `json:"id"`
	OrganisationID OrganisationID `json:"organisation_id"`
	Attributes     Payment۰v0     `json:"attributes"`
}

// CreatePayment۰v0 create a `Payment` in its version 0
func CreatePayment۰v0(
	_ context.Context,
	organisationID OrganisationID,
	attributes Payment۰v0,
) (*Payment, error) {
	ID := aggregate.GenerateID()

	payment := Payment{}
	payment.ID = ID

	paymentCreated := PaymentCreated۰v0{}
	paymentCreated.ID = ID
	paymentCreated.OrganisationID = organisationID
	paymentCreated.Attributes = attributes

	if err := aggregate.RecordChange(&payment, paymentCreated); err != nil {
		return nil, err
	}

	return &payment, nil
}
