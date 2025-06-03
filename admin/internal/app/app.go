// Package app contains all internall app logic.
// It provides App interface with method Run that
// runs full application.
package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"Filmer/admin/config"
	httpController "Filmer/admin/internal/app/controllers/http"
	"Filmer/admin/internal/app/repo"
	"Filmer/admin/internal/app/usecase"
	"Filmer/admin/internal/pkg/logger"
)

var _ App = (*ginApp)(nil)

type App interface {
	Run() error
}

// App implementation
type ginApp struct {
	cfg *config.Config
	log logger.Logger
}

func New(cfg *config.Config) (App, error) {
	return &ginApp{
		cfg: cfg,
		log: logger.NewLogger(os.Stdout, os.Stdout, os.Stderr),
	}, nil
}

// Run starts full application.
func (a ginApp) Run() error {
	router := gin.New()
	// set up base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// set up static
	router.StaticFile("/favicon.ico", "./web/static/img/favicon.ico")
	router.Static("/static", "./web/static")

	// set up handlers
	mainRouterGroup := router.Group("/")
	router.LoadHTMLGlob("./web/template/*")
	// user
	userAPIRepo := repo.NewUserAPIRepo(a.cfg.RestAPI.Host)
	userUC := usecase.NewUserUsecase(userAPIRepo)
	httpController.RegisterUserEndpoints(mainRouterGroup, userUC)

	// create server to run
	server := &http.Server{
		Addr:    ":" + a.cfg.App.Port,
		Handler: router.Handler(),
	}

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.App.KeepAliveTimeout)
	defer cancel()
	// create gracefully shutdown task
	go func() {
		handledSignal := <-quit
		a.log.Infof("Get %q signal. Shutdown server process...", handledSignal.String())
		// shutdown app
		server.Shutdown(shutdownCtx)
	}()

	// start app
	a.log.Infof("Start server on 0.0.0.0:%s...", a.cfg.App.Port)
	err := server.ListenAndServe()
	// unknown error
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("start server: %w", err)
	}

	// wait for gracefully shutdown
	<-shutdownCtx.Done()
	a.log.Infof("Server shutdown successfully!")
	return nil
}
