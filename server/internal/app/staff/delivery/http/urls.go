// Package http contains http router and handlers for personal usecase.
package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

// Personal router.
type StaffRouter struct {
	mwManager              middlewares.MiddlewareManager
	personalHandlerManager *StaffHandlerManager
}

func NewStaffRouter(mwManager middlewares.MiddlewareManager,
	staffHandlerManager *StaffHandlerManager) *StaffRouter {

	return &StaffRouter{
		mwManager:              mwManager,
		personalHandlerManager: staffHandlerManager,
	}
}

// Set routes for handlers in mRouter.personalHandlerManager
func (s StaffRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(s.mwManager.JWTAuth(), s.mwManager.Cache())
	restricted.Get("/full-info/:personID", s.personalHandlerManager.GetPersonInfo())
}
