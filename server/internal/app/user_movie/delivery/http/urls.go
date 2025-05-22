package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

// User movie router
type UserMovieRouter struct {
	mwManager               middlewares.MiddlewareManager
	userMovieHandlerManager *UserMovieHandlerManager
}

// UserMovieRouter constructor
func NewUserMovieRouter(mwManager middlewares.MiddlewareManager, userMovieHandlerManager *UserMovieHandlerManager) *UserMovieRouter {
	return &UserMovieRouter{
		mwManager:               mwManager,
		userMovieHandlerManager: userMovieHandlerManager,
	}
}

// Set routes for handlers in umRouter.userMovieHandlerManager
func (umRouter UserMovieRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(umRouter.mwManager.JWTAuth())
	restricted.Get("/full-info/:kinopoiskID", umRouter.userMovieHandlerManager.GetUserMovie())

	restricted.Get("/stared", umRouter.userMovieHandlerManager.Stared())
	restricted.Get("/want", umRouter.userMovieHandlerManager.Want())
	restricted.Get("/watched", umRouter.userMovieHandlerManager.Watched())

	restricted.Post("/:movieID/star", umRouter.userMovieHandlerManager.Star())
	restricted.Post("/:movieID/unstar", umRouter.userMovieHandlerManager.Unstar())

	restricted.Post("/:movieID/want", umRouter.userMovieHandlerManager.SetWant())
	restricted.Post("/:movieID/watched", umRouter.userMovieHandlerManager.SetWatched())
	restricted.Post("/:movieID/clear", umRouter.userMovieHandlerManager.Clear())
}
