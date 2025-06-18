package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

type MovieRouter struct {
	mwManager           middlewares.MiddlewareManager
	movieHandlerManager *MovieHandlerManager
}

func NewMovieRouter(mwManager middlewares.MiddlewareManager,
	movieHandlerManager *MovieHandlerManager) *MovieRouter {

	return &MovieRouter{
		mwManager:           mwManager,
		movieHandlerManager: movieHandlerManager,
	}
}

// SetRoutes sets routes for handlers in movie handler manager.
func (r MovieRouter) SetRoutes(router fiber.Router) {
	router.Post("/update-movie/:movieID", r.movieHandlerManager.UpdateMovie())

	restricted := router.Use(r.mwManager.JWTAuth(), r.mwManager.Cache())
	restricted.Get("/search", r.movieHandlerManager.SearchFilms())
}
