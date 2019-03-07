package store_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/event/store"
	"github.com/stretchr/testify/assert"
)

func TestPaymentState(t *testing.T) {
	assert.EqualValues(t, &transaction.Payment{}, store.PaymentState())
}

func TestNewTransactionPaymentType۰v0(t *testing.T) {
	aType, err := store.NewTransactionPaymentType۰v0()
	assert.NoError(t, err)

	assert.Equal(t, store.PaymentType, aType.String())
	assert.True(t, aType.IsImplementedBy(&transaction.Payment{}))
}
