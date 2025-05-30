package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

type AuthRouter struct {
	mwManager          middlewares.MiddlewareManager
	authHandlerManager *AuthHandlerManager
}

func NewAuthRouter(mwManager middlewares.MiddlewareManager,
	authHandlerManager *AuthHandlerManager) *AuthRouter {

	return &AuthRouter{
		mwManager:          mwManager,
		authHandlerManager: authHandlerManager,
	}
}

// SetRoutes sets routes for handlers in auth handler manager.
func (r AuthRouter) SetRoutes(router fiber.Router) {
	router.Post("/sign-up", r.authHandlerManager.SignUp())
	router.Post("/login", r.authHandlerManager.Login())

	restricted := router.Use(r.mwManager.JWTAuth())
	restricted.Post("/logout", r.authHandlerManager.Logout())
}
