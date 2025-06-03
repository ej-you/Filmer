// Package http contains HTTP router registrators for
// every entity with its usecases. Also this package
// contains HTTP handlers for handling router endpoints.
package http

import (
	"Filmer/admin/internal/app/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterUserEndpoints defines endpoints for a user usecase and sets up handlers for them.
func RegisterUserEndpoints(router *gin.RouterGroup, userUC usecase.UserUsecase) {
	userHandler := newUserHandler(userUC)
	router.GET("/", userHandler.activityHTML)
}
