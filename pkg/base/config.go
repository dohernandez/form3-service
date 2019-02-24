package base

import (
	logging "github.com/hellofresh/logging-go"
)

const (
	// EnvLive is live environment
	EnvLive = "live"
	// EnvStaging is staging environment
	EnvStaging = "staging"
	// EnvDev is dev environment
	EnvDev = "dev"
)

// Config contains base configuration for any API service
type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME"`
	Port        int    `envconfig:"WEB_PORT" default:"8000"`
	Environment string `envconfig:"ENVIRONMENT" default:"staging"`

	Log logging.LogConfig
}

// LogConfig return Logger Configuration
func (c Config) LogConfig() logging.LogConfig {
	return c.Log
}
