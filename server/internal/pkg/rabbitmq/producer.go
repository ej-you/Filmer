package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

var _ Producer = (*producer)(nil)

// Producer is an interface for publishing messages to RabbitMQ.
type Producer interface {
	PublishText(content []byte) error
}

// producer is a Producer implementation.
type producer struct {
	client    *Client
	queueName string
}

// NewProducer creates new Producer.
func NewProducer(client *Client, queueName string) Producer {
	return &producer{
		client:    client,
		queueName: queueName,
	}
}

// PublishText sends text message to RabbitMQ.
func (p *producer) PublishText(content []byte) error {
	// open channel
	ch, err := p.client.NewChannel()
	if err != nil {
		return fmt.Errorf("open chan: %w", err)
	}
	defer ch.Close()

	// set up queue
	_, err = ch.QueueDeclare(
		p.queueName,
		false, // false means that queue is stored in memory (temporary)
		false, // false means that queue will exist until it is clearly removed
		false, // false means that queue is public (for all connections)
		false, // false means waiting for RabbitMQ confirm that queue is created
		nil,
	)
	if err != nil {
		return fmt.Errorf("declare queue: %w", err)
	}

	// send message
	err = ch.Publish(
		"",
		p.queueName,
		false, // false means no error if message is not match any queue
		false, // (deprecated)
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        content,
		},
	)
	if err != nil {
		return fmt.Errorf("publish message: %w", err)
	}

	return nil
}
