// Package http contains http router and handlers for personal usecase.
package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/middlewares"
)

// Personal router.
type PersonalRouter struct {
	mwManager              middlewares.MiddlewareManager
	personalHandlerManager *PersonalHandlerManager
}

func NewPersonalRouter(mwManager middlewares.MiddlewareManager,
	personalHandlerManager *PersonalHandlerManager) *PersonalRouter {

	return &PersonalRouter{
		mwManager:              mwManager,
		personalHandlerManager: personalHandlerManager,
	}
}

// Set routes for handlers in mRouter.personalHandlerManager
func (p PersonalRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(p.mwManager.JWTAuth())
	restricted.Get("/full-info/:personID", p.personalHandlerManager.GetPersonInfo())
}
