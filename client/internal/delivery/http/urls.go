package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/app/middlewares"
)

// Router to setup client routes
type ClientRouter struct {
	cfg         *config.Config
	mwManager   middlewares.MiddlewareManager
	userHM      *userHandlerManager
	movieHM     *movieHandlerManager
	userMovieHM *userMovieHandlerManager
}

// Router constructor
func NewClientRouter(cfg *config.Config, mwManager middlewares.MiddlewareManager) *ClientRouter {
	return &ClientRouter{
		cfg:         cfg,
		mwManager:   mwManager,
		userHM:      newUserHandlerManager(cfg),
		movieHM:     newMovieHandlerManager(cfg),
		userMovieHM: newUserMovieHandlerManager(cfg),
	}
}

// Main func to setup all of routes
func (r ClientRouter) SetRoutes(router fiber.Router) {
	// main prefix
	appGroup = router.Group(r.cfg.App.PathPrefix)

	appGroup.Get("/", indexGET)

	userGroup := appGroup.Group("/user")
	r.setUserRoutes(userGroup)

	movieGroup := appGroup.Group("/movie")
	r.setMovieRoutes(movieGroup)

	userMovieGroup := appGroup.Group("/user-movie")
	r.setUserMovieRoutes(userMovieGroup)
}

// Setup user subroutes
func (r ClientRouter) setUserRoutes(router fiber.Router) {
	router.Get("/login", r.mwManager.ToProfileIfCookie(), r.userHM.loginGET)
	router.Get("/sign-up", r.mwManager.ToProfileIfCookie(), r.userHM.signUpGET)

	router.Post("/login", r.userHM.loginPOST)
	router.Post("/sign-up", r.userHM.signUpPOST)

	restricted := router.Use(r.mwManager.CookieParser())
	restricted.Post("/logout", r.userHM.logoutPOST)
	restricted.Post("/change-password", r.userHM.changePasswordPOST)

	profile := restricted.Use(r.mwManager.ToLoginIfNoCookie())
	profile.Get("/profile", r.userHM.profileGET)
	profile.Get("/change-password", r.userHM.profileGET)
}

// Setup movie subroutes
func (r ClientRouter) setMovieRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.CookieParser(), r.mwManager.ToLoginIfNoCookie())
	restricted.Get("/search", r.movieHM.searchGET)
}

// Setup user movie subroutes
func (r ClientRouter) setUserMovieRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.CookieParser())
	restrictedWithRedirect := restricted.Use(r.mwManager.ToLoginIfNoCookie())

	restrictedWithRedirect.Get("/info/:kinopoiskID", r.userMovieHM.movieGET)
	restrictedWithRedirect.Get("/stared", r.userMovieHM.staredGET)
	restrictedWithRedirect.Get("/want", r.userMovieHM.wantGET)
	restrictedWithRedirect.Get("/watched", r.userMovieHM.watchedGET)

	restricted.Post("/:movieID/star", r.userMovieHM.starPOST)
	restricted.Post("/:movieID/unstar", r.userMovieHM.unstarPOST)
	restricted.Post("/:movieID/clear", r.userMovieHM.clearPOST)
	restricted.Post("/:movieID/want", r.userMovieHM.wantPOST)
	restricted.Post("/:movieID/watched", r.userMovieHM.watchedPOST)
}
