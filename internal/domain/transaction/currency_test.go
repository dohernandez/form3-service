package transaction_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyValidate(t *testing.T) {
	testCases := []struct {
		scenario string
		currency transaction.Currency
		err      error
	}{
		{
			scenario: "Currency is valid",
			currency: "USD",
		},
		{
			scenario: "Currency is invalid due to is empty",
			currency: "",
			err:      errors.Errorf("cannot be blank"),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			err := tc.currency.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
