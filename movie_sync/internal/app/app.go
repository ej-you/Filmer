// Package app contains all internall app logic.
// It provides App interface with method Run that
// runs full application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"Filmer/movie_sync/config"
	"Filmer/movie_sync/internal/app/adapter/amqp"
	"Filmer/movie_sync/internal/app/repo"
	"Filmer/movie_sync/internal/app/usecase"
	"Filmer/movie_sync/internal/pkg/rabbitmq"
)

var _ App = (*application)(nil)

type App interface {
	Run() error
}

// App implementation
type application struct {
	cfg          *config.Config
	rabbitClient *rabbitmq.Client

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func New(cfg *config.Config) (App, error) {
	// create RabbitMQ client
	rabbitClient, err := rabbitmq.NewClient(cfg.RabbitMQ.ConnURL)
	if err != nil {
		return nil, fmt.Errorf("connect to RabbitMQ: %w", err)
	}
	// init app struct
	application := &application{
		rabbitClient: rabbitClient,
		cfg:          cfg,
	}
	// append main app context
	application.ctx, application.ctxCancel = context.WithCancel(context.Background())
	return application, nil
}

// Run starts full application.
func (a application) Run() error {
	// init repos
	movieAPIRepo := repo.NewMovieAPIRepo(a.cfg.RestAPI.Host)
	// init usecases
	movieUC := usecase.NewMovieUsecase(movieAPIRepo)
	// init adapters
	movieAdapter, err := amqp.NewMovieAdapter(a.rabbitClient, movieUC)
	if err != nil {
		return fmt.Errorf("init movie adapter: %w", err)
	}

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdown := make(chan struct{})
	// create gracefully shutdown task
	go func() {
		handledSignal := <-quit
		logrus.Infof("Get %q signal. Shutdown app...", handledSignal.String())
		a.shutdown()
		shutdown <- struct{}{}
	}()

	// start adapters
	logrus.Info("Start app...")
	if err := movieAdapter.Start(a.ctx); err != nil {
		logrus.Error(err)
		quit <- syscall.SIGTERM
	}

	// wait for gracefully shutdown
	<-shutdown
	logrus.Info("App shutdown successfully!")
	return nil
}

// shutdown gracefully stops running app.
// This method used in Run method for gracefully shutdown.
// After app is stopped it cannot be run again.
func (a application) shutdown() {
	a.ctxCancel()
	if err := a.rabbitClient.Close(); err != nil {
		logrus.Errorf("shutdown: close rabbit client: %v", err)
	}
}
