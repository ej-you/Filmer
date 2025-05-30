package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

type UserRouter struct {
	mwManager          middlewares.MiddlewareManager
	userHandlerManager *UserHandlerManager
}

func NewUserRouter(mwManager middlewares.MiddlewareManager,
	userHandlerManager *UserHandlerManager) *UserRouter {

	return &UserRouter{
		mwManager:          mwManager,
		userHandlerManager: userHandlerManager,
	}
}

// SetRoutes sets routes for handlers in user handler manager.
func (r UserRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.JWTAuth())
	restricted.Post("/change-password", r.userHandlerManager.ChangePassword())
}
