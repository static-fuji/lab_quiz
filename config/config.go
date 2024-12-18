package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"LAB_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"LAB_DB_HOST" envDefault:"lab-db"`
	DBPort     int    `env:"LAB_DB_PORT" envDefault:"3306"`
	DBUser     string `env:"LAB_DB_USER" envDefault:"lab"`
	DBPassword string `env:"LAB_DB_PASSWORD" envDefault:"lab"`
	DBName     string `env:"LAB_DB_DATABASE" envDefault:"lab"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
