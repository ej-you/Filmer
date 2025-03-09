package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
)


// Movie router
type MovieRouter struct {
    mwManager			middlewares.MiddlewareManager
    movieHandlerManager	*MovieHandlerManager
}

// MovieRouter constructor
func NewMovieRouter(mwManager middlewares.MiddlewareManager, movieHandlerManager *MovieHandlerManager) *MovieRouter {
    return &MovieRouter{
    	mwManager: mwManager,
    	movieHandlerManager: movieHandlerManager,
    }
}

// Set routes for handlers in this.movieHandlerManager
func (this MovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(this.mwManager.JWTAuth())
	restricted.Get("/search", this.movieHandlerManager.SearchFilms())
}
