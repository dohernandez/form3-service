package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/dohernandez/form3-service/internal/domain"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/http/handler/transaction/payment"
	"github.com/dohernandez/form3-service/internal/platform/projection/handler/message"
	"github.com/dohernandez/form3-service/pkg/log"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/jmoiron/sqlx"
)

type (
	// PaymentStorage to hold necessary dependencies to manage payment projection
	PaymentStorage struct {
		db    *sqlx.DB
		table string
	}
)

// NewPaymentStorage creates a payment storage
func NewPaymentStorage(db *sqlx.DB, table string) *PaymentStorage {
	return &PaymentStorage{
		db:    db,
		table: table,
	}
}

var _ message.PaymentCreator = new(PaymentStorage)

// Create defines the way to persist a payment in the projection
func (s *PaymentStorage) Create(
	ctx context.Context,
	id aggregate.ID,
	version transaction.Version,
	organisationID transaction.OrganisationID,
	attributes interface{},
) error {
	logger := log.FromContext(ctx)

	query := `INSERT INTO %[1]s (id, version, organisation_id, attributes) VALUES ($1, $2, $3, $4)`
	query = fmt.Sprintf(query, s.table)

	if logger != nil {
		logger.Debugf("exec in transaction sql %s, values %+v", query, []interface{}{
			id,
			version,
			organisationID,
			attributes,
		})
	}

	return execInTransaction(s.db, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, query, id, version, organisationID, attributes)
		if err != nil {
			return err
		}

		return nil
	})
}

var _ message.PaymentBeneficiaryUpdater = new(PaymentStorage)

// UpdateBeneficiary defines the way to update a payment's beneficiary in the projection
func (s *PaymentStorage) UpdateBeneficiary(ctx context.Context, id aggregate.ID, beneficiary transaction.BankAccount) error {
	logger := log.FromContext(ctx)

	jsBeneficiary, err := json.Marshal(beneficiary)
	if err != nil {
		return err
	}

	query := `UPDATE %[1]s 
			  SET attributes = attributes::jsonb || '{"beneficiary_party": %[2]s}'::jsonb
			  WHERE id = $1`
	query = fmt.Sprintf(query, s.table, string(jsBeneficiary))

	if logger != nil {
		logger.Debugf("exec in transaction sql %s, values %+v", query, []interface{}{
			id,
		})
	}

	return execInTransaction(s.db, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}

		return nil
	})
}

var _ message.PaymentDeleter = new(PaymentStorage)

// Delete defines the way to delete the payment in the projection
func (s *PaymentStorage) Delete(ctx context.Context, id aggregate.ID) error {
	logger := log.FromContext(ctx)

	query := `DELETE FROM %[1]s WHERE id = $1`
	query = fmt.Sprintf(query, s.table)

	if logger != nil {
		logger.Debugf("exec in transaction sql %s, values %+v", query, []interface{}{
			id,
		})
	}

	return execInTransaction(s.db, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, query, id)
		if err != nil {
			return err
		}

		return nil
	})
}

var _ payment.FindByID۰v0 = new(PaymentStorage)

// Find defines the way to get the payment from the projection
func (s *PaymentStorage) Find(ctx context.Context, id aggregate.ID) (*transaction.Payment, error) {
	logger := log.FromContext(ctx)

	query := `SELECT * FROM %[1]s WHERE id = $1`
	query = fmt.Sprintf(query, s.table)

	if logger != nil {
		logger.Debugf("exec in transaction sql %s, values %+v", query, []interface{}{
			id,
		})
	}

	var payment transaction.Payment

	err := execInTransaction(s.db, func(tx *sqlx.Tx) error {
		err := tx.GetContext(ctx, &payment, query, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return domain.ErrNotFound
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var attrPayment۰v0 transaction.Payment۰v0

	err = json.Unmarshal(payment.Attributes.([]uint8), &attrPayment۰v0)
	if err != nil {
		return nil, err
	}

	payment.Attributes = attrPayment۰v0

	return &payment, nil
}
