package app

import (
	"github.com/dohernandez/form3-service/pkg/log"
	"github.com/pkg/errors"
)

// NewAppContainer initializes application container
func NewAppContainer(cfg Config) (*Container, error) {
	c := newContainer(cfg)

	logger, err := log.NewLog(cfg.Log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init logger")
	}
	c.WithLogger(logger)

	return c, nil
}
