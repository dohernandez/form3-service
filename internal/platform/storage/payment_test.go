package storage_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/dohernandez/form3-service/internal/domain"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/storage"
	"github.com/dohernandez/form3-service/pkg/test/mocked"
	"github.com/google/uuid"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/stretchr/testify/assert"
)

const (
	table = "table"
)

func TestCreate(t *testing.T) {
	aggregateID := aggregate.GenerateID()
	version := transaction.Version0
	organisationID := transaction.OrganisationID(uuid.New().String())
	attributes := transaction.NewPayment۰v0Mock()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			ID aggregate.ID,
			version transaction.Version,
			organisationID transaction.OrganisationID,
			attributes interface{},
			table string,
		)

		err error
	}{
		{
			scenario: "Create payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				version transaction.Version,
				organisationID transaction.OrganisationID,
				attributes interface{},
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(`^INSERT INTO %[1]s \(.+\) VALUES \(.+\)$`, table)).WithArgs(
					ID,
					version,
					organisationID,
					attributes,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Create payment unsuccessful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				version transaction.Version,
				organisationID transaction.OrganisationID,
				attributes interface{},
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(`^INSERT INTO %[1]s \(.+\) VALUES \(.+\)$`, table)).WithArgs(
					ID,
					version,
					organisationID,
					attributes,
				).WillReturnError(sql.ErrTxDone)

				mock.ExpectRollback()
			},
			err: sql.ErrTxDone,
		},
	}

	dbMock := mocked.NewDBMock(t)
	defer dbMock.Close()

	paymentCreator := storage.NewPaymentStorage(dbMock.SqlxDB, table)
	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			tc.assert(dbMock.Sqlmock, aggregateID, version, organisationID, attributes, table)

			err := paymentCreator.Create(
				context.TODO(),
				aggregateID,
				version,
				organisationID,
				attributes,
			)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}

func TestUpdateBeneficiary(t *testing.T) {
	aggregateID := aggregate.GenerateID()
	beneficiary := transaction.NewBankAccountMock()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			ID aggregate.ID,
			beneficiary transaction.BankAccount,
			table string,
		)

		err error
	}{
		{
			scenario: "Update beneficiary payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				beneficiary transaction.BankAccount,
				table string,
			) {
				jsBeneficiary, err := json.Marshal(beneficiary)
				if err != nil {
					return
				}

				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(
					`^UPDATE %[1]s `+
						`SET attributes = attributes::jsonb \|\| '{"beneficiary_party": %[2]s}'::jsonb `+
						`WHERE id = \$1$`,
					table,
					jsBeneficiary,
				)).WithArgs(
					ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Update beneficiary payment unsuccessful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				beneficiary transaction.BankAccount,
				table string,
			) {
				jsBeneficiary, err := json.Marshal(beneficiary)
				if err != nil {
					return
				}

				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(
					`^UPDATE %[1]s `+
						`SET attributes = attributes::jsonb \|\| '{"beneficiary_party": %[2]s}'::jsonb `+
						`WHERE id = \$1$`,
					table,
					jsBeneficiary,
				)).WithArgs(
					ID,
				).WillReturnError(sql.ErrTxDone)

				mock.ExpectRollback()
			},
			err: sql.ErrTxDone,
		},
	}

	dbMock := mocked.NewDBMock(t)
	defer dbMock.Close()

	paymentBeneficiaryUpdater := storage.NewPaymentStorage(dbMock.SqlxDB, table)
	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			tc.assert(dbMock.Sqlmock, aggregateID, beneficiary, table)

			err := paymentBeneficiaryUpdater.UpdateBeneficiary(
				context.TODO(),
				aggregateID,
				beneficiary,
			)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}

func TestDelete(t *testing.T) {
	aggregateID := aggregate.GenerateID()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			ID aggregate.ID,
			table string,
		)

		err error
	}{
		{
			scenario: "Delete payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(
					`^DELETE FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Delete payment unsuccessful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectExec(fmt.Sprintf(
					`^DELETE FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					ID,
				).WillReturnError(sql.ErrTxDone)

				mock.ExpectRollback()
			},
			err: sql.ErrTxDone,
		},
	}

	dbMock := mocked.NewDBMock(t)
	defer dbMock.Close()

	paymentDeleter := storage.NewPaymentStorage(dbMock.SqlxDB, table)
	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			tc.assert(dbMock.Sqlmock, aggregateID, table)

			err := paymentDeleter.Delete(
				context.TODO(),
				aggregateID,
			)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}

func TestGetPayment(t *testing.T) {
	aggregateID := aggregate.GenerateID()
	version := transaction.Version0
	organisationID := transaction.OrganisationID(uuid.New().String())
	attributes := transaction.NewPayment۰v0Mock()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			ID aggregate.ID,
			table string,
		)

		err error
	}{
		{
			scenario: "Select payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				table string,
			) {
				attrs, err := json.Marshal(attributes)
				assert.NoError(t, err)

				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					ID,
				).WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"version",
					"organisation_id",
					"attributes",
				}).AddRow(
					aggregateID,
					version,
					organisationID,
					attrs,
				))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Select payment unsuccessful, not found",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					ID,
				).WillReturnError(sql.ErrNoRows)

				mock.ExpectRollback()
			},
			err: domain.ErrNotFound,
		},
		{
			scenario: "Select payment unsuccessful, tx error",
			assert: func(
				mock sqlmock.Sqlmock,
				ID aggregate.ID,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					ID,
				).WillReturnError(sql.ErrTxDone)

				mock.ExpectRollback()
			},
			err: sql.ErrTxDone,
		},
	}

	dbMock := mocked.NewDBMock(t)
	defer dbMock.Close()

	paymentFindByID := storage.NewPaymentStorage(dbMock.SqlxDB, table)
	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			tc.assert(dbMock.Sqlmock, aggregateID, table)

			payment, err := paymentFindByID.Find(
				context.TODO(),
				aggregateID,
			)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &transaction.Payment{
					ID:             aggregateID,
					Version:        version,
					OrganisationID: organisationID,
					Attributes:     attributes,
				}, payment)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}
