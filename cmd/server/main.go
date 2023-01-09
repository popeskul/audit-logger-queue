package main

import (
	"github.com/joho/godotenv"
	"github.com/popeskul/audit-logger-queue/pkg/config"
	"github.com/popeskul/audit-logger-queue/pkg/consumer"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error reading env variables from file: %s\n", err.Error())
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error loading env variables: %s\n", err.Error())
	}

	queueConsumer, err := consumer.New(*cfg)
	if err != nil {
		log.Fatalf("failed to initialize queue consumer: %s\n", err.Error())
	}
	defer queueConsumer.Close()

	queueConsumer.Consume()
}
