package beneficiary

import (
	"net/http"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/pkg/event"
	"github.com/dohernandez/form3-service/pkg/http/rest"
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// NewPatchHandler۰v0 creates update beneficiary payment transaction handler
// Handle PATCH /v1/transaction/payments/{id}/beneficiary
func NewPatchHandler۰v0(c interface {
	PaymentEventStore() *event.Store
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		d, err := decodePatchRequest(ctx, r)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrBadRequest(err)))

			return
		}
		form, ok := d.(*PatchRequest۰v0)
		if !ok {
			must.NotFail(render.Render(w, r, rest.ErrInternal(errors.New("wrong request version"))))

			return
		}

		if err := form.Validate(); err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInvalidRequest(err)))

			return
		}

		aggregateRoot, err := c.PaymentEventStore().Get(ctx, form.ID)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrNotFound(err)))

			return
		}

		payment, ok := aggregateRoot.(*transaction.Payment)
		if !ok {
			must.NotFail(render.Render(w, r, rest.ErrInternal(errors.New("mismatch request resource"))))
		}

		err = payment.UpdatePaymentBeneficiary۰v0(ctx, form.Beneficiary)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		err = c.PaymentEventStore().Save(ctx, payment)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		must.NotFail(render.Render(w, r, rest.NoContent))
	}
}
