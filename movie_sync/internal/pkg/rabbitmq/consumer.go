package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var _ Consumer = (*consumer)(nil)

// ConsumeHandler handles every gotten message from RabbitMQ.
type ConsumeHandler func(msg amqp.Delivery)

// Consumer is an interface for getting messages from RabbitMQ.
type Consumer interface {
	Consume(ctx context.Context, handler ConsumeHandler) error
}

// consumer is a Consumer implementation.
type consumer struct {
	client    *Client
	queueName string
}

// NewConsumer creates new Consumer.
func NewConsumer(client *Client, queueName string) (Consumer, error) {
	// open channel (to set up queue)
	channel, err := client.newChannel()
	if err != nil {
		return nil, fmt.Errorf("open chan: %w", err)
	}
	defer channel.Close()

	// set up queue
	_, err = channel.QueueDeclare(
		queueName,
		false, // false means that queue is stored in memory (temporary)
		false, // false means that queue will exist until it is clearly removed
		false, // false means that queue is public (for all connections)
		false, // false means waiting for RabbitMQ confirm that queue is created
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("declare queue: %w", err)
	}

	return &consumer{
		client:    client,
		queueName: queueName,
	}, nil
}

// Consume gets messages from RabbitMQ and handle them with handler.
// This method is blocking. When the ctx will be canceled, consumer will be canceled too.
func (c *consumer) Consume(ctx context.Context, handler ConsumeHandler) error {
	// open channel
	channel, err := c.client.newChannel()
	if err != nil {
		return fmt.Errorf("open chan: %w", err)
	}
	defer channel.Close()

	// create consumer
	messages, err := channel.ConsumeWithContext(
		ctx,
		c.queueName,
		"",
		false, // false means disable messages autoapply
		false, // false means that many consumers can read queue simultaneously
		false, // false means that producer will get its own messages
		false, // false means waiting for producer creating
		nil,
	)
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}

	// handle incoming messages
	for msg := range messages {
		handler(msg)
		msg.Ack(false)
	}
	return nil
}
