package client

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"

	// authHTTP "Filmer/server/internal/auth/delivery/http"
	// movieHTTP "Filmer/server/internal/movie/delivery/http"
	// userMovieHTTP "Filmer/server/internal/user_movie/delivery/http"

	"Filmer/client/config"
	// "Filmer/server/internal/app/middlewares"
	// "Filmer/server/pkg/cache"
	// "Filmer/server/pkg/database"
	// "Filmer/server/pkg/jsonify"
	// "Filmer/server/pkg/logger"
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
	// log     logger.Logger
	// jsonify jsonify.JSONify
}

// Client constructor
func NewClient(cfg *config.Config) Client {
	return &fiberCLient{
		cfg: cfg,
		// log:     logger.NewLogger(),
		// jsonify: jsonify.NewJSONify(),
	}
}

func (s fiberCLient) Run() {
	// app init
	fibertApp := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s v1.0.0", s.cfg.App.Name),
		// ErrorHandler: utils.CustomErrorHandler,
		// JSONEncoder:  s.jsonify.Marshal,
		// JSONDecoder:  s.jsonify.Unmarshal,
		// https://www.f5.com/company/blog/nginx/socket-sharding-nginx-release-1-9-1
		Prefork:      false, // true,
		ServerHeader: s.cfg.App.Name,
	})

	// // DB client init
	// appDB := database.NewCockroachClient(s.cfg, s.log)
	// // cache init
	// appCache := cache.NewCache(s.cfg, s.log)
	// // input data validator init
	// validator := validator.NewValidator()

	// // set up base middlewares
	// mwManager := middlewares.NewMiddlewareManager(s.cfg, appDB, appCache)
	// fibertApp.Use(mwManager.Logger())
	// fibertApp.Use(mwManager.Recover())
	// fibertApp.Use(mwManager.CORS())
	// fibertApp.Use(mwManager.Swagger())

	// // set up handlers
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
		if err := fibertApp.Listen(fmt.Sprintf(":%s", s.cfg.App.Port)); err != nil {
			// s.log.Fatal("[FATAL] failed to start client:", err)
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
		// s.log.Fatal("[FATAL] failed to shutdown client:", err)
		panic(err)
	}
	// s.log.Infof("Client process %d shutdown successfully!", os.Getpid())
	fmt.Printf("Client process %d shutdown successfully!", os.Getpid())
}
