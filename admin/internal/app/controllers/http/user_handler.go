package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Filmer/admin/internal/app/usecase"
	"Filmer/admin/internal/pkg/logger"
)

// Handlers for user usecases.
type userHandler struct {
	userUC usecase.UserUsecase
	log    logger.Logger
}

func newUserHandler(userUC usecase.UserUsecase, log logger.Logger) *userHandler {
	return &userHandler{
		userUC: userUC,
		log:    log,
	}
}

// activityHTML renders HTML for users activity data.
func (g *userHandler) activityHTML(ctx *gin.Context) {
	activity, err := g.userUC.GetUsersActivity()
	if err != nil {
		g.log.Error(err)
	}
	// render html with activity data
	ctx.HTML(http.StatusOK, "user_activity.html", gin.H{
		"usersActivity": activity,
	})
}
