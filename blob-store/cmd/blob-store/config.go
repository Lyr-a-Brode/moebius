package main

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	App   AppOpts
	Store StoreOpts
}

type AppOpts struct {
	Address         string        `env:"ADDRESS" envDefault:":8081"`
	MetricsAddress  string        `env:"METRICS_ADDRESS" envDefault:":9081"`
	EnableDebug     bool          `env:"ENABLE_DEBUG" envDefault:"false"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
}

type StoreOpts struct {
	Type string `env:"STORE_TYPE" envDefault:"disk"`
	Path string `env:"STORE_PATH" envDefault:"/tmp"`
}

func parseConfig() (Config, error) {
	var cfg Config

	opts := env.Options{Prefix: "BS_"}

	if err := env.Parse(&cfg, opts); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
