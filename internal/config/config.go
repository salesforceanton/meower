package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const ENV_PREFIX = "meower"

type Config struct {
	PostgresUsername string `envconfig:"POSTGRES_USERNAME"`
	PostgresName     string `envconfig:"POSTGRES_NAME"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost     string `envconfig:"MEOWER_POSTGRES_HOST"`
	Port             string `envconfig:"PORT"`
	NatsAddress      string `envconfig:"NATS_ADDRESS"`
}

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
