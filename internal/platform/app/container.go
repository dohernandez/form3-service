package app

import (
	"github.com/dohernandez/form3-service/pkg/app"
	"github.com/sirupsen/logrus"
)

// Container contains application resources
type Container struct {
	upstream *app.Container

	cfg Config
}

func newContainer(cfg Config, upstream *app.Container) *Container {
	return &Container{
		upstream: upstream,
		cfg:      cfg,
	}
}

// Cfg returns app-level application configuration
func (c *Container) Cfg() Config {
	return c.cfg
}

// Logger returns app-level logger
func (c *Container) Logger() logrus.FieldLogger {
	return c.upstream.Logger()
}
