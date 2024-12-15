package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("adsboard", *cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
