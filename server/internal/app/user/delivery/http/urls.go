package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

// User router
type UserRouter struct {
	mwManager          middlewares.MiddlewareManager
	userHandlerManager *UserHandlerManager
}

// UserRouter constructor
func NewUserRouter(mwManager middlewares.MiddlewareManager, userHandlerManager *UserHandlerManager) *UserRouter {
	return &UserRouter{
		mwManager:          mwManager,
		userHandlerManager: userHandlerManager,
	}
}

// Set routes for handlers in uRouter.userHandlerManager
func (uRouter UserRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(uRouter.mwManager.JWTAuth())
	restricted.Post("/change-password", uRouter.userHandlerManager.ChangePassword())
}
