package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/app/middlewares"
)

// Router to setup client routes
type ClientRouter struct {
	cfg       *config.Config
	mwManager middlewares.MiddlewareManager
	userHM    *userHandlerManager
	movieHM   *movieHandlerManager
}

// Router constructor
func NewClientRouter(cfg *config.Config, mwManager middlewares.MiddlewareManager) *ClientRouter {
	return &ClientRouter{
		cfg:       cfg,
		mwManager: mwManager,
		userHM:    newUserHandlerManager(cfg),
		movieHM:   newMovieHandlerManager(cfg),
	}
}

// Main func to setup all of routes
func (r ClientRouter) SetRoutes(router fiber.Router) {
	router.Get("/", indexGET)

	userGroup := router.Group("/user")
	r.setUserRoutes(userGroup)

	movieGroup := router.Group("/movie")
	r.setMovieRoutes(movieGroup)
}

// Setup user subroutes
func (r ClientRouter) setUserRoutes(router fiber.Router) {
	router.Get("/login", r.mwManager.ToProfileIfCookie(), r.userHM.loginGET)
	router.Get("/sign-up", r.mwManager.ToProfileIfCookie(), r.userHM.signUpGET)

	router.Post("/login", r.userHM.loginPOST)
	router.Post("/sign-up", r.userHM.signUpPOST)

	restricted := router.Use(r.mwManager.CookieParser())
	restricted.Get("/profile", r.mwManager.ToLoginIfNoCookie(), r.userHM.profileGET)
	restricted.Post("/logout", r.userHM.logoutPOST)
}

// Setup movie subroutes
func (r ClientRouter) setMovieRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.CookieParser(), r.mwManager.ToLoginIfNoCookie())
	restricted.Get("/search", r.movieHM.searchGET)
}
