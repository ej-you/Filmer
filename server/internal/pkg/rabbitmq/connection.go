// Package rabbitmq provides client for connection to RabbitMQ and
// Producer interface to send messages to RabbitMQ.
// AMQP protocol is used for all these actions.
package rabbitmq

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const _closeTimeout = 5 * time.Second // client close timeout

// Client is a client for RabbitMQ via AMQP protocol.
type Client struct {
	conn *amqp.Connection
}

// NewClient creates new RabbitMQ client instance.
func NewClient(connURL string) (*Client, error) {
	conn, err := amqp.Dial(connURL)
	if err != nil {
		return nil, fmt.Errorf("connect to rabbitmq: %w", err)
	}

	return &Client{
		conn: conn,
	}, nil
}

// Close closes RabbitMQ connection.
func (c *Client) Close() error {
	closeDeadline := time.Now().UTC().Add(_closeTimeout)
	err := c.conn.CloseDeadline(closeDeadline)
	if err != nil {
		return fmt.Errorf("close rabbitmq conn: %w", err)
	}
	return nil
}

// NewChannel returns new channel of RabbitMQ connection.
func (c *Client) NewChannel() (*amqp.Channel, error) {
	return c.conn.Channel()
}
