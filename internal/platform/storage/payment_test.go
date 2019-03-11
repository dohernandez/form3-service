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

func TestFindPayment(t *testing.T) {
	aggregateID := aggregate.GenerateID()
	version := transaction.Version0
	organisationID := transaction.OrganisationID(uuid.New().String())
	attributes := transaction.NewPayment۰v0Mock()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			row *transaction.Payment,
			table string,
		)

		err error
	}{
		{
			scenario: "Select payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				row *transaction.Payment,
				table string,
			) {
				attrs, err := json.Marshal(row.Attributes)
				assert.NoError(t, err)

				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					row.ID,
				).WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"version",
					"organisation_id",
					"attributes",
				}).AddRow(
					row.ID,
					row.Version,
					row.OrganisationID,
					attrs,
				))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Select payment unsuccessful, not found",
			assert: func(
				mock sqlmock.Sqlmock,
				row *transaction.Payment,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					row.ID,
				).WillReturnError(sql.ErrNoRows)

				mock.ExpectRollback()
			},
			err: domain.ErrNotFound,
		},
		{
			scenario: "Select payment unsuccessful, tx error",
			assert: func(
				mock sqlmock.Sqlmock,
				row *transaction.Payment,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s WHERE id = \$1$`,
					table,
				)).WithArgs(
					row.ID,
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
			row := transaction.Payment{
				ID:             aggregateID,
				Version:        version,
				OrganisationID: organisationID,
				Attributes:     attributes,
			}

			tc.assert(dbMock.Sqlmock, &row, table)

			payment, err := paymentFindByID.Find(
				context.TODO(),
				aggregateID,
			)

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &row, payment)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}

func TestFindAllPayment(t *testing.T) {
	aggregateID1 := aggregate.GenerateID()
	aggregateID2 := aggregate.GenerateID()

	version := transaction.Version0
	organisationID := transaction.OrganisationID(uuid.New().String())
	attributes1 := transaction.NewPayment۰v0Mock()
	attributes2 := transaction.NewPayment۰v0Mock()

	testCases := []struct {
		scenario string

		assert func(
			mock sqlmock.Sqlmock,
			rows []*transaction.Payment,
			table string,
		)

		err error
	}{
		{
			scenario: "Select all payment successful",
			assert: func(
				mock sqlmock.Sqlmock,
				rows []*transaction.Payment,
				table string,
			) {
				attrs1, err := json.Marshal(rows[0].Attributes)
				assert.NoError(t, err)

				attrs2, err := json.Marshal(rows[1].Attributes)
				assert.NoError(t, err)

				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s$`,
					table,
				)).WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"version",
					"organisation_id",
					"attributes",
				}).AddRow(
					rows[0].ID,
					rows[0].Version,
					rows[0].OrganisationID,
					attrs1,
				).AddRow(
					rows[1].ID,
					rows[1].Version,
					rows[1].OrganisationID,
					attrs2,
				))

				mock.ExpectCommit()
			},
		},
		{
			scenario: "Select all payment unsuccessful, tx error",
			assert: func(
				mock sqlmock.Sqlmock,
				rows []*transaction.Payment,
				table string,
			) {
				mock.ExpectBegin()

				// #nosec G201
				mock.ExpectQuery(fmt.Sprintf(
					`^SELECT \* FROM %[1]s$`,
					table,
				)).WillReturnError(sql.ErrTxDone)

				mock.ExpectRollback()
			},
			err: sql.ErrTxDone,
		},
	}

	dbMock := mocked.NewDBMock(t)
	defer dbMock.Close()

	paymentFindAll := storage.NewPaymentStorage(dbMock.SqlxDB, table)
	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			rows := []*transaction.Payment{
				{
					ID:             aggregateID1,
					OrganisationID: organisationID,
					Version:        version,
					Attributes:     attributes1,
				},
				{
					ID:             aggregateID2,
					OrganisationID: organisationID,
					Version:        version,
					Attributes:     attributes2,
				},
			}

			tc.assert(dbMock.Sqlmock, rows, table)

			payments, err := paymentFindAll.FindAll(context.TODO())

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, rows, payments)
			}

			err = dbMock.Sqlmock.ExpectationsWereMet()
			assert.NoErrorf(t, err, "there were unfulfilled expectations: %s", err)
		})
	}
}
