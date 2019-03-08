package payment

import (
	"net/http"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/pkg/event"
	"github.com/dohernandez/form3-service/pkg/http/rest"
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// PostRequest۰v0 represents the request inputs considered to create a transaction payment
type PostRequest۰v0 struct {
	OrganisationID transaction.OrganisationID `json:"organisation_id"`
	Attributes     transaction.Payment۰v0     `json:"attributes"`
}

// NewPostHandler۰v0 creates a create payment transaction handler
// Handle POST /v1/transaction/payments
func NewPostHandler۰v0(c interface {
	PaymentEventStore() *event.Store
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		d, err := decodePostRequest(ctx, r)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrBadRequest(err)))

			return
		}
		form, ok := d.(*PostRequest۰v0)
		if !ok {
			must.NotFail(render.Render(w, r, rest.ErrInternal(errors.New("wrong request version"))))

			return
		}

		if err := form.Validate(); err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInvalidRequest(err)))

			return
		}

		payment, err := transaction.CreatePayment۰v0(ctx, form.OrganisationID, form.Attributes)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		err = c.PaymentEventStore().Save(ctx, payment)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		w.WriteHeader(http.StatusCreated)

		render.JSON(w, r, encodeToResponse(ctx, payment))
	}
}
