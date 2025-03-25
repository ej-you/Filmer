package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"

	httpDelivery "Filmer/client/internal/delivery/http"
	// movieHTTP "Filmer/server/internal/movie/delivery/http"
	// userMovieHTTP "Filmer/server/internal/user_movie/delivery/http"

	"Filmer/client/config"
	"Filmer/client/internal/app/middlewares"
	"Filmer/client/internal/pkg/logger"
	// "Filmer/server/pkg/cache"
	// "Filmer/server/pkg/database"
	// "Filmer/server/pkg/jsonify"
	// "Filmer/server/pkg/utils"
	// "Filmer/server/pkg/validator"
)

// Client interface
type Client interface {
	Run()
}

// Fiber client
type fiberCLient struct {
	cfg *config.Config
	log logger.Logger
	// jsonify jsonify.JSONify
}

// Client constructor
func NewClient(cfg *config.Config) Client {
	return &fiberCLient{
		cfg: cfg,
		log: logger.NewLogger(),
		// jsonify: jsonify.NewJSONify(),
	}
}

func (c fiberCLient) Run() {
	// template engine
	engine := django.New("./web/template", ".html")

	// app init
	fibertApp := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s v1.0.0", c.cfg.App.Name),
		// ErrorHandler: utils.CustomErrorHandler,
		ServerHeader: c.cfg.App.Name,
		Views:        engine,
	})

	// set up base middlewares
	mwManager := middlewares.NewMiddlewareManager(c.cfg)
	fibertApp.Use(mwManager.Logger())
	fibertApp.Use(mwManager.Recover())

	// set up static
	fibertApp.Static("/favicon.ico", "./web/static/img/favicon.ico")
	fibertApp.Static("/static", "./web/static")
	// set up handlers
	router := httpDelivery.NewClientRouter(c.cfg, mwManager)
	router.SetRoutes(fibertApp)

	// apiV1 := fibertApp.Group("/api/v1")
	// // auth
	// authHandlerManager := authHTTP.NewAuthHandlerManager(s.cfg, appDB, appCache, validator)
	// authRouter := authHTTP.NewAuthRouter(mwManager, authHandlerManager)
	// authRouter.SetRoutes(apiV1.Group("/user"))
	// // movie
	// movieHandlerManager := movieHTTP.NewMovieHandlerManager(s.cfg, s.jsonify, s.log, appDB, appCache, validator)
	// movieRouter := movieHTTP.NewMovieRouter(mwManager, movieHandlerManager)
	// movieRouter.SetRoutes(apiV1.Group("/kinopoisk/films"))
	// // user movie
	// userMovieHandlerManager := userMovieHTTP.NewUserMovieHandlerManager(s.cfg, s.jsonify, s.log, appDB, appCache, validator)
	// userMovieRouter := userMovieHTTP.NewUserMovieRouter(mwManager, userMovieHandlerManager)
	// userMovieRouter.SetRoutes(apiV1.Group("/films"))

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
