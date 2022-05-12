package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
	"time"
)

const (
	ReportIntervalDefault = time.Second * 10
	PollIntervalDefault   = time.Second * 2
	DefaultHost           = "localhost:8080"
)

type Config struct {
	Address        string        `env:"ADDRESS"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
}

func NewAgentConfig() *Config {
	return Config{}.Init()
}

func (cfg Config) Init() *Config {
	flag.StringVar(&cfg.Address, "a", DefaultHost, "agent address")
	flag.DurationVar(&cfg.PollInterval, "p", PollIntervalDefault, "report interval")
	flag.DurationVar(&cfg.ReportInterval, "r", ReportIntervalDefault, "report interval")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	log.Println(cfg, "cfg")
	return &cfg
}
