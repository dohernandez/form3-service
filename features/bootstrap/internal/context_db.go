package internal

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/dohernandez/form3-service/pkg/test/feature"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DBContext struct to hold necessary dependencies
type DBContext struct {
	feature.DBContext

	paymentIDs map[string]string
}

// RegisterDBContext is the place to truncate database, run background givens and check then should be stored
func RegisterDBContext(s *godog.Suite, db *sqlx.DB) *DBContext {
	c := DBContext{
		DBContext: feature.DBContext{
			DB: db,
			Tables: []string{
				"transaction_payment",
				"transaction_projections",
				"events_transaction_stream",
			},
		},
		paymentIDs: map[string]string{},
	}

	s.BeforeScenario(func(_ interface{}) {
		c.CleanUpDB()

		// #nosec G201
		query := `INSERT INTO transaction_projections (no, name, position, state, locked) VALUES ($1, $2, $3, $4, $5)`
		_, err := c.DB.Exec(query, []interface{}{1, "transaction_payment", 0, "{}", false}...)
		if err != nil {
			return
		}
	})

	s.Step(`^the following payment\(s\) should be stored in the table "([^"]*)"$`, c.theFollowingPaymentsShouldBeStoredInTheTable)
	s.Step(`^the following payment state should be stored in the table "([^"]*)"$`, c.theFollowingPaymentStateShouldBeStoredInTheTable)
	s.Step(`^that the following payment state\(s\) are stored in the table "([^"]*)"$`, c.thatTheFollowingPaymentStatesAreStoredInTheTable)
	s.Step(`^that the following payment\(s\) are stored in the table "([^"]*)"$`, c.thatTheFollowingPaymentsAreStoredInTheTable)

	return &c
}

func (c *DBContext) theFollowingPaymentsShouldBeStoredInTheTable(table string, data *gherkin.DataTable) error {
	cIds, err := c.theFollowingElementsShouldBeStoredInTheTable("ID", "id", table, data, c.paymentValueWhereBuilder)
	if err != nil {
		return err
	}

	for k, id := range cIds {
		c.paymentIDs[k] = id
	}

	return nil
}

func (c *DBContext) theFollowingElementsShouldBeStoredInTheTable(idPlaceHolder, selectCol, table string, data *gherkin.DataTable, valueWhereBuilder feature.ValueWhereBuilder) (cIds map[string]string, err error) {
	var (
		skipColumnsFromCondition []string
		found                    bool
		ids                      []string
		idColIndex               int
	)

	if idPlaceHolder != "" {
		found, idColIndex = c.findColumnIndex(data, idPlaceHolder)
		if found {
			skipColumnsFromCondition = append(skipColumnsFromCondition, idPlaceHolder)
		}
	}

	ids, err = c.RunExistData(selectCol, table, data, valueWhereBuilder, skipColumnsFromCondition)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find expected data")
	}

	if found {
		cIds = make(map[string]string)

		for k, row := range data.Rows[1:] {
			v := row.Cells[idColIndex].Value

			cIds[v] = ids[k]
		}
	}

	return cIds, nil
}

func (c *DBContext) findColumnIndex(data *gherkin.DataTable, colName string) (bool, int) {
	var (
		found    bool
		colIndex int
	)

	for k, cell := range data.Rows[0].Cells {
		if cell.Value == colName {
			colIndex = k
			found = true

			break
		}
	}

	return found, colIndex
}

func (c *DBContext) paymentValueWhereBuilder(col, v string, position int) (string, interface{}) {
	switch col {
	case "id":
		id := uuid.MustParse(v)

		return fmt.Sprintf("%s = $%d", col, position), id

	case "attributes":
		var value transaction.Payment۰v0

		err := json.Unmarshal([]byte(v), &value)
		must.NotFail(err)

		return fmt.Sprintf("%s @> $%d", col, position), value
	}

	if v == "" {
		return fmt.Sprintf("%s IS NULL", col), nil
	}

	return fmt.Sprintf("%s = $%d", col, position), v
}

func (c *DBContext) theFollowingPaymentStateShouldBeStoredInTheTable(table string, data *gherkin.DataTable) error {
	_, err := c.theFollowingElementsShouldBeStoredInTheTable("", "event_id", table, data, c.paymentStateValueWhereBuilder)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBContext) paymentStateValueWhereBuilder(col, v string, position int) (string, interface{}) {
	if col == "METADATA" || col == "metadata" {
		if col == "METADATA" {
			var value struct {
				ID      string `json:"_aggregate_id"`
				Type    string `json:"_aggregate_type"`
				Version int    `json:"_aggregate_version"`
			}

			err := json.Unmarshal([]byte(v), &value)
			must.NotFail(err)

			value.ID = c.paymentIDs[value.ID]

			vb, err := json.Marshal(value)
			must.NotFail(err)

			return fmt.Sprintf("%s @> $%d", strings.ToLower(col), position), vb
		}

		return fmt.Sprintf("%s @> $%d", strings.ToLower(col), position), v
	}

	if v == "" {
		return fmt.Sprintf("%s IS NULL", col), nil
	}

	return fmt.Sprintf("%s = $%d", col, position), v
}

func (c *DBContext) thatTheFollowingPaymentStatesAreStoredInTheTable(table string, data *gherkin.DataTable) error {
	err := c.RunStoreData(table, data, nil)
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	return nil
}

func (c *DBContext) thatTheFollowingPaymentsAreStoredInTheTable(table string, data *gherkin.DataTable) error {
	cIds, err := c.theFollowingElementsShouldBeStoredInTheTable("", "id", table, data, c.paymentValueWhereBuilder)
	if err != nil {
		return err
	}

	for k, id := range cIds {
		c.paymentIDs[k] = id
	}

	return nil
}
