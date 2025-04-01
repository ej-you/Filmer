package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"

	httpDelivery "Filmer/client/internal/delivery/http"

	"Filmer/client/config"
	"Filmer/client/internal/app/middlewares"
	"Filmer/client/internal/pkg/logger"
)

// Client interface
type Client interface {
	Run()
}

// Fiber client
type fiberCLient struct {
	cfg *config.Config
	log logger.Logger
}

// Client constructor
func NewClient(cfg *config.Config) Client {
	return &fiberCLient{
		cfg: cfg,
		log: logger.NewLogger(),
	}
}

func (c fiberCLient) Run() {
	// template engine
	engine := django.New("./web/template", ".html")

	// app init
	fibertApp := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("%s v1.0.0", c.cfg.App.Name),
		ErrorHandler: httpDelivery.CustomErrorHandler,
		ServerHeader: c.cfg.App.Name,
		Views:        engine,
	})

	// set up base middlewares
	mwManager := middlewares.NewMiddlewareManager(c.cfg)
	fibertApp.Use(mwManager.Logger())
	fibertApp.Use(mwManager.Recover())
	fibertApp.Use(mwManager.Compression())

	// set up static
	fibertApp.Static("/favicon.ico", "./web/static/img/favicon.ico")
	fibertApp.Static("/static", "./web/static")
	// set up handlers
	router := httpDelivery.NewClientRouter(c.cfg, mwManager)
	router.SetRoutes(fibertApp)

	// start client
	go func() {
		if err := fibertApp.Listen(fmt.Sprintf(":%s", c.cfg.App.Port)); err != nil {
			c.log.Fatal("[FATAL] failed to start client:", err)
			panic(err)
		}
	}()

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quit

	// shutdown client
	if err := fibertApp.Shutdown(); err != nil {
		c.log.Fatal("[FATAL] failed to shutdown client:", err)
		panic(err)
	}
	c.log.Info("Client shutdown successfully!")
}
