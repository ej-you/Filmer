package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
)

// Movie router
type MovieRouter struct {
	mwManager           middlewares.MiddlewareManager
	movieHandlerManager *MovieHandlerManager
}

// MovieRouter constructor
func NewMovieRouter(mwManager middlewares.MiddlewareManager, movieHandlerManager *MovieHandlerManager) *MovieRouter {
	return &MovieRouter{
		mwManager:           mwManager,
		movieHandlerManager: movieHandlerManager,
	}
}

// Set routes for handlers in mRouter.movieHandlerManager
func (mRouter MovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(mRouter.mwManager.JWTAuth())
	restricted.Get("/search", mRouter.movieHandlerManager.SearchFilms())
}
