package config

import "time"

type Config struct {
	Address       string        `env:"ADDRESS" envDefault:"localhost:8080"`
	StoreInterval time.Duration `env:"STORE_INTERVAL" envDefault:"0s"`
	FileName      string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	ShouldRestore bool          `env:"RESTORE" envDefault:"true"`
}
