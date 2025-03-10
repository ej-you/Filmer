package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
)


// User movie router
type UserMovieRouter struct {
    mwManager				middlewares.MiddlewareManager
    userMovieHandlerManager	*UserMovieHandlerManager
}

// UserMovieRouter constructor
func NewUserMovieRouter(mwManager middlewares.MiddlewareManager, userMovieHandlerManager *UserMovieHandlerManager) *UserMovieRouter {
    return &UserMovieRouter{
    	mwManager: mwManager,
    	userMovieHandlerManager: userMovieHandlerManager,
    }
}

// Set routes for handlers in this.userMovieHandlerManager
func (this UserMovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(this.mwManager.JWTAuth())
	restricted.Get("/full-info/:kinopoiskID", this.userMovieHandlerManager.GetUserMovie())

	restricted.Get("/stared", this.userMovieHandlerManager.Stared())
	restricted.Get("/want", this.userMovieHandlerManager.Want())
	restricted.Get("/watched", this.userMovieHandlerManager.Watched())
	
	restricted.Post("/:movieID/star", this.userMovieHandlerManager.Star())
	restricted.Post("/:movieID/unstar", this.userMovieHandlerManager.Unstar())
	
	restricted.Post("/:movieID/want", this.userMovieHandlerManager.SetWant())
	restricted.Post("/:movieID/watched", this.userMovieHandlerManager.SetWatched())
	restricted.Post("/:movieID/clear", this.userMovieHandlerManager.Clear())
}
