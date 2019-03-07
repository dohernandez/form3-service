package app

import (
	"github.com/dohernandez/form3-service/pkg/app"
	"github.com/dohernandez/form3-service/pkg/event"
)

// Container contains application resources
type Container struct {
	*app.Container

	cfg               Config
	paymentEventStore *event.Store
}

func newContainer(cfg Config, upstream *app.Container) *Container {
	return &Container{
		Container: upstream,
		cfg:       cfg,
	}
}

// Cfg returns app-level application configuration
// nolint:unused
func (c *Container) Cfg() Config {
	return c.cfg
}

// WithPaymentEventStore sets eventStore.EventStore instance
func (c *Container) WithPaymentEventStore(paymentEventStore *event.Store) *Container {
	c.paymentEventStore = paymentEventStore

	return c
}

// PaymentEventStore returns app-level event.Store instance
func (c *Container) PaymentEventStore() *event.Store {
	return c.paymentEventStore
}
