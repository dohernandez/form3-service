package transaction

import validation "github.com/go-ozzo/ozzo-validation"

// BankAccount represents a person bank account
type BankAccount struct {
	Person
	Account
	AccountName       AccountName       `json:"account_name"`
	AccountNumberCode AccountNumberCode `json:"account_number_code"`
}

// Validate validates BankAccount
func (ba BankAccount) Validate() error {
	return validation.ValidateStruct(
		&ba,
		// AccountName using AccountName validation rules
		validation.Field(&ba.AccountName),
		// AccountNumberCode using AccountNumberCode validation rules
		validation.Field(&ba.AccountNumberCode),
		// Validate Person using Person validation rules
		validation.Field(&ba.Person),
		// Validate Account using Account validation rules
		validation.Field(&ba.Account),
	)
}

// AccountName represents an account name
type AccountName string

// Validate validates AccountName cannot be empty
func (an AccountName) Validate() error {
	v := string(an)

	return validation.Validate(&v, validation.Required)
}

// AccountNumberCode represents an account number code
type AccountNumberCode string

// Validate validates AccountNumberCode cannot be empty
func (anc AccountNumberCode) Validate() error {
	v := string(anc)

	return validation.Validate(&v, validation.Required)
}

type (
	// Account represents a back account
	Account struct {
		Bank
		Number AccountNumber `json:"account_number"`
		Type   AccountType   `json:"account_type,omitempty"`
	}

	// AccountType represents an account type
	AccountType int64
)

// Validate validates Account
func (a Account) Validate() error {
	return validation.ValidateStruct(
		&a,
		// Number cannot using AccountNumber validation rules
		validation.Field(&a.Number),
		// Validate Bank using Bank validation rules
		validation.Field(&a.Bank),
	)
}

// AccountNumber represents an account number
type AccountNumber string

// Validate validates AccountNumber cannot be empty
func (an AccountNumber) Validate() error {
	v := string(an)

	return validation.Validate(v, validation.Required)
}

// Bank represents a bank
type Bank struct {
	ID     BankID     `json:"bank_id,string"`
	IDCode BankIDCode `json:"bank_id_code"`
}

// Validate validates Bank
func (b Bank) Validate() error {
	return validation.ValidateStruct(
		&b,
		// ID cannot be empty
		validation.Field(&b.ID, validation.Required.Error("cannot be zero")),
		// IDCode cannot be empty
		validation.Field(&b.IDCode, validation.Required),
	)
}

// BankID represents bank id
type BankID int64

// Validate validates BankID cannot be empty
func (bi BankID) Validate() error {
	v := int64(bi)

	return validation.Validate(v, validation.Required.Error("cannot be zero"))
}

// BankIDCode represents bank id code
type BankIDCode string

// Validate validates BankIDCode cannot be empty
func (bic BankIDCode) Validate() error {
	v := string(bic)

	return validation.Validate(v, validation.Required)
}
