package transaction

import validation "github.com/go-ozzo/ozzo-validation"

type (
	// ChargesInformation represents charge information
	ChargesInformation struct {
		BearerCode              BearerCode     `json:"bearer_code"`
		SenderCharges           []SenderCharge `json:"sender_charges"`
		ReceiverChargesAmount   float64        `json:"receiver_charges_amount,string"`
		ReceiverChargesCurrency Currency       `json:"receiver_charges_currency"`
	}
)

// Validate validates ChargesInformation
func (ci ChargesInformation) Validate() error {
	return validation.ValidateStruct(
		&ci,
		// BearerCode using BearerCode validation rules
		validation.Field(&ci.BearerCode),
		// ReceiverChargesAmount cannot be empty
		validation.Field(&ci.ReceiverChargesAmount, validation.Required.Error("cannot be zero")),
		// Validate ReceiverChargesCurrency using Currency validation rules
		validation.Field(&ci.ReceiverChargesCurrency),
		// SenderCharges cannot be empty
		validation.Field(&ci.SenderCharges, validation.Required.Error("cannot be empty")),
		// Validate SenderCharge using SenderCharge validation rules
		validation.Field(&ci.SenderCharges),
	)
}

// BearerCode represents charge information barcode
type BearerCode string

// Validate validates BearerCode cannot be empty
func (an BearerCode) Validate() error {
	v := string(an)

	return validation.Validate(&v, validation.Required)
}

// SenderCharge represents charge information sender charge
type SenderCharge struct {
	Amount   float64  `json:"amount,string"`
	Currency Currency `json:"currency"`
}

// Validate validates SenderCharge
func (sc SenderCharge) Validate() error {
	return validation.ValidateStruct(
		&sc,
		// Amount cannot be empty
		validation.Field(&sc.Amount, validation.Required.Error("cannot be zero")),
		// Validate Currency using Currency validation rules
		validation.Field(&sc.Currency),
	)
}
