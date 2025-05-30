package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

type UserMovieRouter struct {
	mwManager               middlewares.MiddlewareManager
	userMovieHandlerManager *UserMovieHandlerManager
}

func NewUserMovieRouter(mwManager middlewares.MiddlewareManager,
	userMovieHandlerManager *UserMovieHandlerManager) *UserMovieRouter {

	return &UserMovieRouter{
		mwManager:               mwManager,
		userMovieHandlerManager: userMovieHandlerManager,
	}
}

// SetRoutes sets routes for handlers in user-movie handler manager.
func (r UserMovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.JWTAuth())
	restricted.Get("/full-info/:kinopoiskID", r.userMovieHandlerManager.GetUserMovie())

	restricted.Get("/stared", r.userMovieHandlerManager.Stared())
	restricted.Get("/want", r.userMovieHandlerManager.Want())
	restricted.Get("/watched", r.userMovieHandlerManager.Watched())

	restricted.Post("/:movieID/star", r.userMovieHandlerManager.Star())
	restricted.Post("/:movieID/unstar", r.userMovieHandlerManager.Unstar())

	restricted.Post("/:movieID/want", r.userMovieHandlerManager.SetWant())
	restricted.Post("/:movieID/watched", r.userMovieHandlerManager.SetWatched())
	restricted.Post("/:movieID/clear", r.userMovieHandlerManager.Clear())
}
