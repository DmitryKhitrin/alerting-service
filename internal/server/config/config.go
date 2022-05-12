package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
	"time"
)

const (
	DefaultAddress       = "localhost:8080"
	ShouldRestore        = true
	StoreIntervalDefault = time.Second * 300
	FileNameDefault      = "/tmp/devops-metrics-db.json"
)

type Config struct {
	Address       string        `env:"ADDRESS"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	FileName      string        `env:"STORE_FILE"`
	ShouldRestore bool          `env:"RESTORE"`
}

func NewSeverConfig() *Config {
	return Config{}.Init()

}

func (cfg Config) Init() *Config {
	flag.StringVar(&cfg.Address, "a", DefaultAddress, "server address")
	flag.BoolVar(&cfg.ShouldRestore, "r", ShouldRestore, "restore data")
	flag.DurationVar(&cfg.StoreInterval, "i", StoreIntervalDefault, "store interval")
	flag.StringVar(&cfg.FileName, "f", FileNameDefault, "store file")
	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Println(cfg, "мой конфижек с ошибкой")
		panic(err)
	}

	log.Println(cfg, "мой конфижек")

	return &cfg
}
