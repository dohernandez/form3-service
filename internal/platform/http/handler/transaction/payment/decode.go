package payment

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

// PostRequest۰v0 represents the request inputs considered to create a payment
type PostRequest۰v0 struct {
	OrganisationID transaction.OrganisationID `json:"organisation_id"`
	Attributes     transaction.Payment۰v0     `json:"attributes"`
}

// Validate validates PatchRequest۰v0
func (pr PostRequest۰v0) Validate() error {
	return validation.ValidateStruct(
		&pr,
		// ID cannot be empty and must be uuid
		validation.Field(&pr.OrganisationID, validation.Required, is.UUID),
		// Attributes using Payment۰v0 validation rules
		validation.Field(&pr.Attributes),
	)
}

// decodePostRequest decode request into Form
func decodePostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var form PostRequest۰v0

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		return nil, errors.Wrap(err, "Decoding post json body")
	}

	return &form, nil
}

// DeleteRequest represents the request inputs considered to delete a payment
type DeleteRequest struct {
	ID aggregate.ID
}

// Validate validates DeleteRequest
func (pr DeleteRequest) Validate() error {
	return validation.ValidateStruct(
		&pr,
		// ID cannot be empty and must be uuid
		validation.Field(&pr.ID, validation.Required, is.UUID),
	)
}

// decodeDeleteRequest decode request into Form
func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var form DeleteRequest

	ID, err := aggregate.IDFromString(chi.URLParam(r, "id"))
	if err != nil {
		return nil, err
	}

	form.ID = ID

	return &form, nil
}
