package consumer

import (
	"fmt"
	"log"

	"github.com/popeskul/audit-logger-queue/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func New(cfg config.Config) (*Consumer, error) {
	connectionStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Queue.User, cfg.Queue.Password, cfg.Queue.Host, cfg.Queue.Port)

	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		logrus.Fatalf("failed to connect to RabbitMQ: %s", err)
		return nil, fmt.Errorf("failed to connect to rabbitmq %w/n", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		logrus.Fatalf("failed to open a channel: %s", err)
		return nil, fmt.Errorf("failed to open rabbitmq channel: %w/n", err)
	}

	queue, err := channel.QueueDeclare(
		cfg.Queue.Topic, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		logrus.Fatalf("failed to declare a queue: %s", err)
		return nil, fmt.Errorf("failed to declare rabbitmq queue: %w/n", err)
	}

	return &Consumer{
		channel: channel,
		queue:   &queue,
	}, nil
}

func (q *Consumer) Close() error {
	return q.channel.Close()
}

func (q *Consumer) Consume() {
	messages, err := q.channel.Consume(
		q.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("failed to consume %s", err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range messages {
			logrus.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
