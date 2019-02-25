package http

import (
	"github.com/dohernandez/form3-service/internal/platform/app"
	"github.com/dohernandez/form3-service/pkg/http/router"
	"github.com/go-chi/chi"
)

func NewRouter(c *app.Container) chi.Router {
	r := router.NewRouter(c.Container)

	return r
}
