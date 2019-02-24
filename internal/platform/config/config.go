package config

import (
	"github.com/dohernandez/form3-service/pkg/base"
	"github.com/kelseyhightower/envconfig"
)

// Specification contains structured configuration variables.
type Specification struct {
	base.Config
}

// LoadEnv load config variables into baseConfig.
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	config.ServiceName = "form3-service"

	return &config, err
}
