// Package http contains http router and handlers for staff usecase.
package http

import (
	fiber "github.com/gofiber/fiber/v2"

	"Filmer/server/internal/app/server/middlewares"
)

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

// SetRoutes sets routes for handlers in staff handler manager.
func (r StaffRouter) SetRoutes(router fiber.Router) {
	restricted := router.Use(r.mwManager.JWTAuth(), r.mwManager.Cache())
	restricted.Get("/full-info/:personID", r.personalHandlerManager.GetPersonInfo())
}
