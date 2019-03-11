package payment

import (
	"net/http"

	"github.com/dohernandez/form3-service/internal/domain"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/internal/platform"
	"github.com/dohernandez/form3-service/pkg/event"
	"github.com/dohernandez/form3-service/pkg/http/rest"
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// NewDeleteHandler creates delete a payment handler
// Handle DELETE /v1/transaction/payments/{id}
func NewDeleteHandler(c interface {
	PaymentEventStore() *event.Store
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// decoding request
		d, err := decodeDeleteRequest(ctx, r)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrBadRequest(err)))

			return
		}
		form, ok := d.(*DeleteRequest)
		if !ok {
			must.NotFail(render.Render(w, r, rest.ErrInternal(platform.ErrWrongRequestVersion)))

			return
		}

		// validating request
		if err := form.Validate(); err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInvalidRequest(err)))

			return
		}

		// finding payment root
		aggregateRoot, err := c.PaymentEventStore().Get(ctx, form.ID)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrNotFound(err)))

			return
		}

		payment, ok := aggregateRoot.(*transaction.Payment)
		if !ok {
			must.NotFail(render.Render(w, r, rest.ErrInternal(platform.ErrMismatchRequest)))
		}

		// checking if payment is already deleted
		if payment.Deleted {
			must.NotFail(render.Render(w, r, rest.ErrNotFound(errors.Wrap(err, domain.ErrNotFound.Error()))))
		}

		// delete payment
		err = payment.DeletePayment(ctx)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		// save payment state
		err = c.PaymentEventStore().Save(ctx, payment)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		must.NotFail(render.Render(w, r, rest.NoContent))
	}
}
