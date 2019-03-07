package http

import (
	"net/http"

	"github.com/dohernandez/form3-service/internal/platform/app"
	"github.com/dohernandez/form3-service/internal/platform/http/handler/transaction/payment"
	"github.com/dohernandez/form3-service/pkg/http/router"
	"github.com/go-chi/chi"
)

// NewRouter creates an instance of router filled with handlers and docs
func NewRouter(c *app.Container) chi.Router {
	r := router.NewRouter(c.Container)

	logger := c.Logger()

	paymentURI := "/v1/transaction/payments"
	r.Method(http.MethodPost, paymentURI, payment.NewPostHandler۰v0(c))
	logger.Debugf("added `%s %s` route", http.MethodPost, paymentURI)

	return r
}
