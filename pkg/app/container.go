package app

import "github.com/sirupsen/logrus"

// Container is a basic application container
type Container struct {
	cfg Config

	logger logrus.FieldLogger
}

func newContainer(cfg Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

// Cfg returns app-level base configuration
func (c *Container) Cfg(cfg Config) Config {
	return c.cfg
}

// WithLogger sets logger instance
func (c *Container) WithLogger(logger logrus.FieldLogger) {
	c.logger = logger
}

// Logger returns app-level logger
func (c *Container) Logger() logrus.FieldLogger {
	if c.logger == nil {
		c.logger = logrus.New()
	}

	return c.logger
}
