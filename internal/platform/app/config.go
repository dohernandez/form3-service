package app

import (
	"github.com/dohernandez/form3-service/pkg/app"
	"github.com/kelseyhightower/envconfig"
)

// Config contains structured configuration variables.
type Config struct {
	app.Config
}

// LoadEnv load env variables into Config.
func LoadEnv() (Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}
