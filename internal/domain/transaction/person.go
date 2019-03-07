package transaction

import validation "github.com/go-ozzo/ozzo-validation"

// Person represents a person
type Person struct {
	Name    Name    `json:"name"`
	Address Address `json:"address"`
}

// Validate validates Person
func (p Person) Validate() error {
	return validation.ValidateStruct(
		&p,
		// Name using Name validation rules
		validation.Field(&p.Name),
		// Address using Address validation rules
		validation.Field(&p.Address),
	)
}

// Name represents a person name
type Name string

// Validate validates Name cannot be empty
func (n Name) Validate() error {
	v := string(n)

	return validation.Validate(&v, validation.Required)
}

// Address represents a person address
type Address string

// Validate validates Address cannot be empty
func (a Address) Validate() error {
	v := string(a)

	return validation.Validate(&v, validation.Required)
}
