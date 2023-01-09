package main

import (
	"github.com/joho/godotenv"
	"github.com/popeskul/audit-logger-queue/pkg/client"
	"github.com/popeskul/audit-logger-queue/pkg/config"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error reading env variables from file: %s\n", err)
	}

	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err)
	}

	queueLogger, err := client.New(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize queue client: %s\n", err)
	}
	defer queueLogger.Close()

	// example of how to use the client
	// send a message to the queue
	err = queueLogger.Produce("hello world")
	if err != nil {
		logrus.Fatalf("failed to produce message: %s\n", err)
	}
}
