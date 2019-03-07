package transaction

import validation "github.com/go-ozzo/ozzo-validation"

// Currency represents currency
type Currency string

// Validate validates Currency cannot be empty
func (c Currency) Validate() error {
	v := string(c)

	return validation.Validate(&v, validation.Required)
}
