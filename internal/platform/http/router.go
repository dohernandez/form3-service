package http

import (
	"net/http"

	"github.com/dohernandez/form3-service/internal/platform/app"
	"github.com/dohernandez/form3-service/internal/platform/http/handler/transaction/payment"
	"github.com/dohernandez/form3-service/internal/platform/http/handler/transaction/payment/beneficiary"
	"github.com/dohernandez/form3-service/pkg/http/router"
	"github.com/go-chi/chi"
)

// NewRouter creates an instance of router filled with handlers and docs
func NewRouter(c *app.Container) chi.Router {
	r := router.NewRouter(c.Container)

	logger := c.Logger()

	paymentCreateURI := "/v1/transaction/payments"
	r.Method(http.MethodPost, paymentCreateURI, payment.NewPostHandler۰v0(c))
	logger.Debugf("added `%s %s` route", http.MethodPost, paymentCreateURI)

	paymentBeneficiaryURI := "/v1/transaction/payments/{id}/beneficiary"
	r.Method(http.MethodPatch, paymentBeneficiaryURI, beneficiary.NewPatchHandler۰v0(c))
	logger.Debugf("added `%s %s` route", http.MethodPatch, paymentBeneficiaryURI)

	paymentDeleteURI := "/v1/transaction/payments/{id}"
	r.Method(http.MethodDelete, paymentDeleteURI, payment.NewDeleteHandler(c))
	logger.Debugf("added `%s %s` route", http.MethodDelete, paymentDeleteURI)

	paymentGetURI := "/v1/transaction/payments/{id}"
	r.Method(http.MethodGet, paymentGetURI, payment.NewGetHandler(c))
	logger.Debugf("added `%s %s` route", http.MethodGet, paymentGetURI)

	return r
}
