// Package http contains HTTP router registrators for
// every entity with its usecases. Also this package
// contains HTTP handlers for handling router endpoints.
package http

import (
	"github.com/gin-gonic/gin"

	"Filmer/admin/internal/app/usecase"
	"Filmer/admin/internal/pkg/logger"
)

// RegisterUserEndpoints defines endpoints for a user usecase and sets up handlers for them.
func RegisterUserEndpoints(router *gin.RouterGroup,
	userUC usecase.UserUsecase, log logger.Logger) {

	userHandler := newUserHandler(userUC, log)
	router.GET("/", userHandler.activityHTML)
}
