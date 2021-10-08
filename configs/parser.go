package configs

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// EnvParser gathers configs from the environment.
func EnvParser() (Config, error) {
	var cfg Config
	if err := envconfig.Process("MYMARIADB", &cfg.MariaDB); err != nil {
		return Config{}, errors.Wrap(err, "error processing mariadb configs")
	}

	if err := envconfig.Process("MYSERVER", &cfg.Server); err != nil {
		return Config{}, errors.Wrap(err, "error processing server configs")
	}

	return cfg, nil
}
