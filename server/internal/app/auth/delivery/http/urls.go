package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

// Auth router
type AuthRouter struct {
	mwManager          middlewares.MiddlewareManager
	authHandlerManager *AuthHandlerManager
}

// AuthRouter constructor
func NewAuthRouter(mwManager middlewares.MiddlewareManager, authHandlerManager *AuthHandlerManager) *AuthRouter {
	return &AuthRouter{
		mwManager:          mwManager,
		authHandlerManager: authHandlerManager,
	}
}

// Set routes for handlers in aRouter.authHandlerManager
func (aRouter AuthRouter) SetRoutes(router fiber.Router) {
	router.Post("/sign-up", aRouter.authHandlerManager.SignUp())
	router.Post("/login", aRouter.authHandlerManager.Login())

	restricted := router.Use(aRouter.mwManager.JWTAuth())
	restricted.Post("/logout", aRouter.authHandlerManager.Logout())
}
