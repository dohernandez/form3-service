package beneficiary

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/pkg/errors"
)

// PatchRequest۰v0 represents the data submitted in the request useful to handle the request
type PatchRequest۰v0 struct {
	ID          aggregate.ID
	Beneficiary transaction.BankAccount
}

// Validate validates PatchRequest۰v0
func (pr PatchRequest۰v0) Validate() error {
	return validation.ValidateStruct(
		&pr,
		// ID cannot be empty and must be uuid
		validation.Field(&pr.ID, validation.Required, is.UUID),
		// Beneficiary using BankAccount validation rules
		validation.Field(&pr.Beneficiary),
	)
}

// decodePatchRequest decode request into Form
func decodePatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var form PatchRequest۰v0

	var beneficiary transaction.BankAccount

	err := json.NewDecoder(r.Body).Decode(&beneficiary)
	if err != nil {
		return nil, errors.Wrap(err, "Decoding post json body")
	}

	form.Beneficiary = beneficiary

	ID, err := aggregate.IDFromString(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}

	form.ID = ID

	return &form, nil
}
