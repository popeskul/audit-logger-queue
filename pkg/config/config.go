package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Queue Queue
}

type Queue struct {
	User     string `envconfig:"QUEUE_USER" required:"true"`
	Password string `envconfig:"QUEUE_PASSWORD" required:"true"`
	Host     string `envconfig:"QUEUE_HOST" required:"true"`
	Port     int    `envconfig:"QUEUE_PORT" required:"true"`
	Topic    string `envconfig:"QUEUE_TOPIC" required:"true"`
}

func New() (*Config, error) {
	cfg := Config{}

	if err := envconfig.Process("queue", &cfg.Queue); err != nil {
		return nil, err
	}

	return &cfg, nil
}
