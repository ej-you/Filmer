package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
)


// Auth router
type AuthRouter struct {
    mwManager			middlewares.MiddlewareManager
    authHandlerManager	*AuthHandlerManager
}

// AuthRouter constructor
func NewAuthRouter(mwManager middlewares.MiddlewareManager, authHandlerManager *AuthHandlerManager) *AuthRouter {
    return &AuthRouter{
    	mwManager: mwManager,
    	authHandlerManager: authHandlerManager,
    }
}

// Set routes for handlers in this.authHandlerManager
func (this AuthRouter) SetRoutes(router fiber.Router) {
	router.Post("/sign-up", this.authHandlerManager.SignUp())
	router.Post("/login", this.authHandlerManager.Login())

	restricted := router.Use(this.mwManager.JWTAuth())
	restricted.Post("/logout", this.authHandlerManager.Logout())
}
