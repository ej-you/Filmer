// Package server contains server interface and middleware manager interface for server.
package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"Filmer/server/config"
	authHTTP "Filmer/server/internal/app/auth/delivery/http"
	movieHTTP "Filmer/server/internal/app/movie/delivery/http"
	"Filmer/server/internal/app/server/middlewares"
	staffHTTP "Filmer/server/internal/app/staff/delivery/http"
	userHTTP "Filmer/server/internal/app/user/delivery/http"
	userMovieHTTP "Filmer/server/internal/app/usermovie/delivery/http"
	"Filmer/server/internal/pkg/cache"
	"Filmer/server/internal/pkg/database"
	"Filmer/server/internal/pkg/errhandler"
	"Filmer/server/internal/pkg/jsonify"
	"Filmer/server/internal/pkg/logger"
	"Filmer/server/internal/pkg/validator"
)

var _ Server = (*fiberServer)(nil)

// Server interface.
type Server interface {
	Run() error
}

// Fiber server.
type fiberServer struct {
	cfg          *config.Config
	log          logger.Logger
	jsonify      jsonify.JSONify
	validate     validator.Validator
	dbStorage    *gorm.DB
	cacheStorage cache.Storage
}

// Server constructor.
func New(cfg *config.Config) (Server, error) {
	log := logger.NewLogger(cfg.LogOutput.Info, cfg.LogOutput.Info, cfg.LogOutput.Error)

	// DB storage init
	dbStorage, err := database.New(cfg.Database.ConnString,
		database.WithLogger(log),
		database.WithWarnLogLevel(),
		database.WithDisableColorful(),
		database.WithIgnoreNotFound())
	if err != nil {
		return nil, err
	}
	// cache storage init
	cacheStorage, err := cache.NewStorage(cfg.Cache.ConnString, log)
	if err != nil {
		return nil, err
	}

	return &fiberServer{
		cfg:          cfg,
		log:          log,
		jsonify:      jsonify.NewJSONify(),
		dbStorage:    dbStorage,
		cacheStorage: cacheStorage,
		validate:     validator.New(),
	}, nil
}

//	@title			Filmer API
//	@version		1.0.0
//	@description	This is a Filmer API for Kinopoisk API and DB

//	@license.name	MIT Licence
//	@license.url	https://github.com/ej-you/Filmer/blob/master/LICENCE

//	@host		127.0.0.1:3000
//	@basePath	/api/v1
//	@schemes	http

//	@accept						json
//	@produce					json
//	@query.collection.format	multi

// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
// @description				JWT security accessToken. Please, add it in the format "Bearer {AccessToken}" to authorize your requests.
func (s *fiberServer) Run() error {
	// app init
	fiberApp := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("%s v1.0.0", s.cfg.App.Name),
		ErrorHandler: errhandler.CustomErrorHandler,
		JSONEncoder:  s.jsonify.Marshal,
		JSONDecoder:  s.jsonify.Unmarshal,
		// https://www.f5.com/company/blog/nginx/socket-sharding-nginx-release-1-9-1
		Prefork:      false,
		ServerHeader: s.cfg.App.Name,
	})

	// set up base middlewares
	mwManager := middlewares.NewMiddlewareManager(s.cfg, s.dbStorage, s.cacheStorage)
	fiberApp.Use(mwManager.Logger())
	fiberApp.Use(mwManager.Recover())
	fiberApp.Use(mwManager.CORS())
	fiberApp.Use(mwManager.Swagger())

	// set up handlers
	apiV1 := fiberApp.Group("/api/v1")
	// auth
	authHandlerManager := authHTTP.NewAuthHandlerManager(s.cfg,
		s.dbStorage, s.cacheStorage, s.validate)
	authRouter := authHTTP.NewAuthRouter(mwManager, authHandlerManager)
	authRouter.SetRoutes(apiV1.Group("/auth"))
	// movie
	movieHandlerManager := movieHTTP.NewMovieHandlerManager(s.cfg, s.jsonify, s.log,
		s.dbStorage, s.cacheStorage, s.validate)
	movieRouter := movieHTTP.NewMovieRouter(mwManager, movieHandlerManager)
	movieRouter.SetRoutes(apiV1.Group("/kinopoisk/films"))
	// user movie
	userMovieHandlerManager := userMovieHTTP.NewUserMovieHandlerManager(s.cfg, s.jsonify, s.log,
		s.dbStorage, s.cacheStorage, s.validate)
	userMovieRouter := userMovieHTTP.NewUserMovieRouter(mwManager, userMovieHandlerManager)
	userMovieRouter.SetRoutes(apiV1.Group("/films"))
	// user
	userHandlerManager := userHTTP.NewUserHandlerManager(s.cfg, s.dbStorage, s.validate)
	userRouter := userHTTP.NewUserRouter(mwManager, userHandlerManager)
	userRouter.SetRoutes(apiV1.Group("/user"))
	// staff
	staffHandlerManager := staffHTTP.NewStaffHandlerManager(s.cfg, s.jsonify, s.log, s.validate)
	staffRouter := staffHTTP.NewStaffRouter(mwManager, staffHandlerManager)
	staffRouter.SetRoutes(apiV1.Group("/staff"))

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdownDone := make(chan struct{})
	// create gracefully shutdown task
	go func() {
		handledSignal := <-quit
		s.log.Infof("Get %q signal. Shutdown %d server process...",
			handledSignal.String(), os.Getpid())
		// shutdown app
		fiberApp.ShutdownWithTimeout(s.cfg.App.KeepAliveTimeout)
		shutdownDone <- struct{}{}
	}()

	// start app
	if err := fiberApp.Listen(fmt.Sprintf(":%s", s.cfg.App.Port)); err != nil {
		return fmt.Errorf("start app: %w", err)
	}

	// wait for gracefully shutdown
	<-shutdownDone
	s.log.Infof("Server process %d shutdown successfully!", os.Getpid())
	return nil
}
