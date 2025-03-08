package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"

	authHTTP "Filmer/server/internal/auth/delivery/http"

	"Filmer/server/internal/app/middlewares"
	"Filmer/server/pkg/cache"
	"Filmer/server/pkg/database"
	"Filmer/server/pkg/jsonify"
	"Filmer/server/pkg/logger"
	"Filmer/server/pkg/utils"
	"Filmer/server/pkg/validator"
	"Filmer/server/config"
)

// Server interface
type Server interface {
	Run()
}


// Fiber server
type fiberServer struct {
	cfg			*config.Config
	log			logger.Logger
	jsonify		jsonify.JSONify
}

// Server constructor
func NewServer(cfg *config.Config) Server {
	return &fiberServer{
		cfg: cfg,
		log: logger.NewLogger(),
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

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
//	@description				JWT security accessToken. Please, add it in the format "Bearer {AccessToken}" to authorize your requests.
func (this fiberServer) Run() {
	// app init
	fibertApp := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s v1.0.0", this.cfg.App.Name),
		ErrorHandler: utils.CustomErrorHandler,
		JSONEncoder: this.jsonify.Marshal,
		JSONDecoder: this.jsonify.Unmarshal,
		// https://www.f5.com/company/blog/nginx/socket-sharding-nginx-release-1-9-1
		Prefork: false, //true,
		ServerHeader: this.cfg.App.Name,
	})

	// DB client init
	appDB := database.NewCockroachClient(this.cfg, this.log)
	// cache init
	appCache := cache.NewCache(this.cfg, this.log)
	// input data validator init
	validator := validator.NewValidator()

	// set up base middlewares
	mwManager := middlewares.NewMiddlewareManager(this.cfg, appDB, appCache)
	fibertApp.Use(mwManager.Logger())
	fibertApp.Use(mwManager.Recover())
	fibertApp.Use(mwManager.CORS())
	fibertApp.Use(mwManager.Swagger())

	// set up handlers
	apiV1 := fibertApp.Group("/api/v1")
	// auth
	authHandlerManager := authHTTP.NewAuthHandlerManager(this.cfg, appDB, appCache, validator)
	authRouter := authHTTP.NewAuthRouter(mwManager, authHandlerManager)
	authRouter.SetRoutes(apiV1.Group("/user"))

	// start server
	go func() {
		if err := fibertApp.Listen(fmt.Sprintf(":%s", this.cfg.App.Port)); err != nil {
			this.log.Fatal("[FATAL] failed to start server:", err)
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
		this.log.Fatal("[FATAL] failed to shutdown server:", err)
	}
	this.log.Infof("Server process %d shutdown successfully!", os.Getpid())
}
