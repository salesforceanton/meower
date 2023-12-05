package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const ENV_PREFIX = "meower"

type Config struct {
	PostgresUsername  string `envconfig:"POSTGRES_USERNAME"`
	PostgresName      string `envconfig:"POSTGRES_NAME"`
	PostgresPassword  string `envconfig:"POSTGRES_PASSWORD"`
	PostgresHost      string `envconfig:"POSTGRES_HOST"`
	PostgresPort      string `envconfig:"POSTGRES_PORT"`
	Port              string `envconfig:"PORT"`
	NatsPort          string `envconfig:"NATS_PORT"`
	NatsHost          string `envconfig:"NATS_HOST"`
	ElasticsearchHost string `envconfig:"ELASTICSEARCH_HOST"`
	ElasticsearchPort string `envconfig:"ELASTICSEARCH_PORT"`
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
