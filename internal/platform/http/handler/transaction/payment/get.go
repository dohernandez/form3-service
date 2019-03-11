package payment

import (
	"context"
	"net/http"

	"github.com/dohernandez/form3-service/internal/domain"
	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/dohernandez/form3-service/pkg/http/rest"
	"github.com/dohernandez/form3-service/pkg/log"
	"github.com/dohernandez/form3-service/pkg/must"
	"github.com/go-chi/render"
	"github.com/hellofresh/goengine/aggregate"
	"github.com/sirupsen/logrus"
)

// FindByID۰v0 defines the way to find the payment by id
type FindByID۰v0 interface {
	Find(ctx context.Context, ID aggregate.ID) (*transaction.Payment, error)
}

// NewGetHandler creates delete a payment handler
// Handle GET /v1/transaction/payments/{id}
func NewGetHandler(c interface {
	PaymentFindByID۰v0() FindByID۰v0
	Logger() *logrus.Logger
}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// decoding request
		form, err := decodeGetRequest(ctx, r)
		if err != nil {
			must.NotFail(render.Render(w, r, rest.ErrBadRequest(err)))

			return
		}

		// validating request
		if err := form.Validate(); err != nil {
			must.NotFail(render.Render(w, r, rest.ErrInvalidRequest(err)))

			return
		}

		// read payment
		ctx = log.ToContext(ctx, c.Logger())
		payment, err := c.PaymentFindByID۰v0().Find(ctx, form.ID)
		if err != nil {
			if err == domain.ErrNotFound {
				must.NotFail(render.Render(w, r, rest.ErrNotFound(err)))

				return
			}

			must.NotFail(render.Render(w, r, rest.ErrInternal(err)))

			return
		}

		render.JSON(w, r, encodeToResponse(ctx, payment))
	}
}
