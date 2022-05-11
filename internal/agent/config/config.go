package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"time"
)

const (
	ReportIntervalDefault = time.Second * 10
	PollIntervalDefault   = time.Second * 2
)

type Config struct {
	Address        string        `env:"ADDRESS" envDefault:"localhost:8080"`
	PollInterval   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
}

func NewAgentConfig() *Config {
	return Config{}.Init()
}

func (cfg Config) Init() *Config {

	flag.StringVar(&cfg.Address, "a", "localhost:8080", "agent address")
	flag.DurationVar(&cfg.PollInterval, "p", PollIntervalDefault, "report interval")
	flag.DurationVar(&cfg.ReportInterval, "r", ReportIntervalDefault, "report interval")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
