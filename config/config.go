package config

import (
	"github.com/cerminan/libs/config"
)

type Config struct {
  Host string `default:":50051"`
}

func DefaultConfig() (Config, error) {
  var err error

  var cfg Config
  err = config.Init(&cfg)
  if err != nil {
    return Config{}, nil
  }

  return cfg, nil
}

func (c *Config) LoadEnvar() error {
  var err error
  
  err = config.LoadEnvar(c)
  if err != nil {
    return err
  }

  return nil
}
