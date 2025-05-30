package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

// Movie router
type MovieRouter struct {
	mwManager           middlewares.MiddlewareManager
	movieHandlerManager *MovieHandlerManager
}

// MovieRouter constructor
func NewMovieRouter(mwManager middlewares.MiddlewareManager,
	movieHandlerManager *MovieHandlerManager) *MovieRouter {

	return &MovieRouter{
		mwManager:           mwManager,
		movieHandlerManager: movieHandlerManager,
	}
}

// Set routes for handlers in mRouter.movieHandlerManager
func (m MovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(m.mwManager.JWTAuth(), m.mwManager.Cache())
	restricted.Get("/search", m.movieHandlerManager.SearchFilms())
}
