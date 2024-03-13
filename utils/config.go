package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceMode string `envconfig:"SERVICE_MODE"`
	LogLevel    string `envconfig:"LOG_LEVEL"`

	DBHost string `envconfig:"DB_HOST"`
	DBPort string `envconfig:"DB_PORT"`
	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`

	CMCApiKey   string `envconfig:"CMC_API_KEY"`
	MainnetType string `envconfig:"MAINNET_TYPE"`
}

var cfg Config

func GetConfig() *Config {
	if (Config{}) != cfg {
		return &cfg
	}

	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
