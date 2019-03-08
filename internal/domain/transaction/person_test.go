package transaction_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	testCases := []struct {
		scenario   string
		personFunc func(transaction.Person) transaction.Person
		err        error
	}{
		{
			scenario: "Person is valid",
			personFunc: func(person transaction.Person) transaction.Person {
				return person
			},
		},
		{
			scenario: "Person is invalid due to empty name",
			personFunc: func(person transaction.Person) transaction.Person {
				person.Name = ""

				return person
			},
			err: errors.Errorf("name: cannot be blank."),
		},
		{
			scenario: "Person is invalid due to empty address",
			personFunc: func(person transaction.Person) transaction.Person {
				person.Address = ""

				return person
			},
			err: errors.Errorf("address: cannot be blank."),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			person := tc.personFunc(transaction.NewPersonMock())

			err := person.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
