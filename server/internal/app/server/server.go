package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"

	authHTTP "Filmer/server/internal/auth/delivery/http"
	movieHTTP "Filmer/server/internal/movie/delivery/http"
	personalHTTP "Filmer/server/internal/personal/delivery/http"
	userHTTP "Filmer/server/internal/user/delivery/http"
	userMovieHTTP "Filmer/server/internal/user_movie/delivery/http"

	"Filmer/server/config"
	"Filmer/server/internal/app/middlewares"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/database"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/logger"
	"Filmer/server/pkg/utils"
	"Filmer/server/pkg/validator"
)

// Server interface
type Server interface {
	Run()
}

// Fiber server
type fiberServer struct {
	cfg     *config.Config
	log     logger.Logger
	jsonify jsonify.JSONify
}

// Server constructor
func NewServer(cfg *config.Config) Server {
	return &fiberServer{
		cfg:     cfg,
		log:     logger.NewLogger(cfg),
		jsonify: jsonify.NewJSONify(),
	}
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
func (s fiberServer) Run() {
	// app init
	fibertApp := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("%s v1.0.0", s.cfg.App.Name),
		ErrorHandler: utils.CustomErrorHandler,
		JSONEncoder:  s.jsonify.Marshal,
		JSONDecoder:  s.jsonify.Unmarshal,
		// https://www.f5.com/company/blog/nginx/socket-sharding-nginx-release-1-9-1
		Prefork:      false,
		ServerHeader: s.cfg.App.Name,
	})

	// DB client init
	appDB := database.NewCockroachClient(s.cfg, s.log)
	// cache init
	appCache := cache.NewCache(s.cfg, s.log)
	// input data validator init
	validator := validator.NewValidator()

	// set up base middlewares
	mwManager := middlewares.NewMiddlewareManager(s.cfg, appDB, appCache)
	fibertApp.Use(mwManager.Logger())
	fibertApp.Use(mwManager.Recover())
	fibertApp.Use(mwManager.CORS())
	fibertApp.Use(mwManager.Swagger())

	// set up handlers
	apiV1 := fibertApp.Group("/api/v1")
	// auth
	authHandlerManager := authHTTP.NewAuthHandlerManager(s.cfg, appDB, appCache, validator)
	authRouter := authHTTP.NewAuthRouter(mwManager, authHandlerManager)
	authRouter.SetRoutes(apiV1.Group("/auth"))
	// movie
	movieHandlerManager := movieHTTP.NewMovieHandlerManager(s.cfg, s.jsonify, s.log, appDB, appCache, validator)
	movieRouter := movieHTTP.NewMovieRouter(mwManager, movieHandlerManager)
	movieRouter.SetRoutes(apiV1.Group("/kinopoisk/films"))
	// user movie
	userMovieHandlerManager := userMovieHTTP.NewUserMovieHandlerManager(s.cfg, s.jsonify, s.log, appDB, appCache, validator)
	userMovieRouter := userMovieHTTP.NewUserMovieRouter(mwManager, userMovieHandlerManager)
	userMovieRouter.SetRoutes(apiV1.Group("/films"))
	// user
	userHandlerManager := userHTTP.NewUserHandlerManager(s.cfg, appDB, validator)
	userRouter := userHTTP.NewUserRouter(mwManager, userHandlerManager)
	userRouter.SetRoutes(apiV1.Group("/user"))
	// personal
	personalHandlerManager := personalHTTP.NewPersonalHandlerManager(s.cfg, s.jsonify, s.log, appCache, validator)
	personalRouter := personalHTTP.NewPersonalRouter(mwManager, personalHandlerManager)
	personalRouter.SetRoutes(apiV1.Group("/personal"))

	// start server
	go func() {
		if err := fibertApp.Listen(fmt.Sprintf(":%s", s.cfg.App.Port)); err != nil {
			s.log.Fatal("[FATAL] failed to start server:", err)
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

	// shutdown server
	if err := fibertApp.Shutdown(); err != nil {
		s.log.Fatal("[FATAL] failed to shutdown server:", err)
	}
	s.log.Infof("Server process %d shutdown successfully!", os.Getpid())
}
