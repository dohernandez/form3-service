package message

import (
	"context"
	"time"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/hellofresh/goengine"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/hellofresh/goengine/metadata"
)

// NewCallbackPaymentCreatorMock creates a callback mock for tests
// nolint:unused
func NewCallbackPaymentCreatorMock(createFunc func(
	ID aggregate.ID,
	version transaction.Version,
	organisationID transaction.OrganisationID,
	attributes interface{},
) error) PaymentCreator {
	return &paymentCreatorMock{
		createFunc: createFunc,
	}
}

// nolint:unused
type paymentCreatorMock struct {
	createFunc func(
		id aggregate.ID,
		version transaction.Version,
		organisationID transaction.OrganisationID,
		attributes interface{},
	) error
}

func (m *paymentCreatorMock) Create(
	_ context.Context,
	id aggregate.ID,
	version transaction.Version,
	organisationID transaction.OrganisationID,
	attributes interface{},
) error {
	return m.createFunc(id, version, organisationID, attributes)
}

// NewMessageMock creates a goengine.Message mock for tests
// nolint:unused
func NewMessageMock(payloadFunc func() interface{}) goengine.Message {
	return &messageMock{
		payloadFunc: payloadFunc,
	}
}

// nolint:unused
type messageMock struct {
	payloadFunc func() interface{}
}

func (m *messageMock) UUID() goengine.UUID {
	panic("should not be called")
}

func (m *messageMock) CreatedAt() time.Time {
	panic("should not be called")
}

func (m *messageMock) Payload() interface{} {
	return m.payloadFunc()
}

func (m *messageMock) Metadata() metadata.Metadata {
	panic("should not be called")
}

func (m *messageMock) WithMetadata(key string, value interface{}) goengine.Message {
	panic("should not be called")
}

// NewCallbackPaymentBeneficiaryUpdaterMock creates a callback mock for tests
// nolint:unused
func NewCallbackPaymentBeneficiaryUpdaterMock(beneficiaryUpdaterFunc func(
	ID aggregate.ID,
	beneficiary transaction.BankAccount,
) error) PaymentBeneficiaryUpdater {
	return &paymentBeneficiaryUpdaterMock{
		beneficiaryUpdaterFunc: beneficiaryUpdaterFunc,
	}
}

// nolint:unused
type paymentBeneficiaryUpdaterMock struct {
	beneficiaryUpdaterFunc func(ID aggregate.ID, beneficiary transaction.BankAccount) error
}

func (m *paymentBeneficiaryUpdaterMock) UpdateBeneficiary(ctx context.Context, id aggregate.ID, beneficiary transaction.BankAccount) error {
	return m.beneficiaryUpdaterFunc(id, beneficiary)
}

// NewCallbackPaymentDeleterMock creates a callback mock for tests
// nolint:unused
func NewCallbackPaymentDeleterMock(deleterFunc func(ID aggregate.ID) error) PaymentDeleter {
	return &paymentDeleterMock{
		deleterFunc: deleterFunc,
	}
}

// nolint:unused
type paymentDeleterMock struct {
	deleterFunc func(ID aggregate.ID) error
}

func (m *paymentDeleterMock) Delete(ctx context.Context, id aggregate.ID) error {
	return m.deleterFunc(id)
}
