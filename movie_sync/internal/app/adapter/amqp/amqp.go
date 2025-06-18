// Package amqp contains adapters for RabbitMQ with handlers for them.
package amqp

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"

	"Filmer/movie_sync/internal/app/entity"
	"Filmer/movie_sync/internal/app/usecase"
	"Filmer/movie_sync/internal/pkg/rabbitmq"
)

const _queueName = "movie.update" // name of RabbitMQ queue for movies updates

// MovieAdapter is an adapter for movies updates.
type MovieAdapter struct {
	consumer rabbitmq.Consumer
	movieUC  usecase.MovieUsecase
}

func NewMovieAdapter(rabbitClient *rabbitmq.Client,
	movieUC usecase.MovieUsecase) (*MovieAdapter, error) {

	consumer, err := rabbitmq.NewConsumer(rabbitClient, _queueName)
	if err != nil {
		return nil, fmt.Errorf("init consumer: %w", err)
	}
	return &MovieAdapter{
		consumer: consumer,
		movieUC:  movieUC,
	}, nil
}

// Start starts RabbitMQ consumer of movies to update.
// This method is blocking. Use context with cancel to stop adapter work.
func (a *MovieAdapter) Start(ctx context.Context) error {
	err := a.consumer.Consume(ctx, a.handler())
	if err != nil {
		return fmt.Errorf("start movie adapter: %w", err)
	}
	return nil
}

// handler returns consumer handler for RabbitMQ consumer of movies to update.
func (a *MovieAdapter) handler() rabbitmq.ConsumeHandler {
	return func(msg amqp.Delivery) {
		logrus.Info("New message: movie to update")
		// parse movie id from message
		movie, err := entity.NewMovie(msg.Body)
		if err != nil {
			logrus.Errorf("Create movie instance: %v", err)
			return
		}
		// update movie
		if err := a.movieUC.FullUpdate(movie); err != nil {
			logrus.Errorf("Handle movie %s update: %v", movie.ID(), err)
			return
		}
		logrus.Infof("Movie %s was updated successfully!", movie.ID())
	}
}
