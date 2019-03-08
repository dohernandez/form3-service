package payment

import (
	"context"
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
)

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
