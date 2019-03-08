package storage_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform/storage"
	"github.com/dohernandez/form3-service/pkg/test/mocked"
	"github.com/google/uuid"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	aggregateID := aggregate.GenerateID()
	version := transaction.Version0
	organisationID := transaction.OrganisationID(uuid.New().String())
	attributes := transaction.NewPayment۰v0Mock()
	table := "table"

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
			scenario: "Create payment successfully",
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
			scenario: "Failure create payment",
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
