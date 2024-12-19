package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host  string `default:"0.0.0.0"`
	Port  int    `default:"8080"`
	Debug bool   `default:"true"`
	App   string `default:"coupons"`
}

func New() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
