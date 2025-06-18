// Package amqp contains adapters for RabbitMQ with handlers for them.
package amqp

import (
	"fmt"

	"github.com/google/uuid"

	"Filmer/server/internal/pkg/rabbitmq"
)

const _queueName = "movie.update" // name of RabbitMQ queue for movies updates

// MovieAdapter is an adapter for movies updates.
type MovieAdapter struct {
	producer rabbitmq.Producer
}

func NewMovieAdapter(rabbitClient *rabbitmq.Client) (*MovieAdapter, error) {
	producer, err := rabbitmq.NewProducer(rabbitClient, _queueName)
	if err != nil {
		return nil, fmt.Errorf("init producer: %w", err)
	}
	return &MovieAdapter{
		producer: producer,
	}, nil
}

// SendID sends movie UUID to RabbitMQ.
// Sent message will be used by consumer for update movie info.
func (a *MovieAdapter) SendID(movieID uuid.UUID) error {
	// use [:] to correct convert uuid to byte slice
	err := a.producer.PublishText(movieID[:])
	if err != nil {
		return fmt.Errorf("send movie id: %w", err)
	}
	return nil
}
