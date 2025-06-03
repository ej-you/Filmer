package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Filmer/admin/internal/app/usecase"
)

// Handlers for user usecases.
type userHandler struct {
	userUC usecase.UserUsecase
}

func newUserHandler(userUC usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUC: userUC,
	}
}

// activityHTML renders HTML for users activity data.
func (g *userHandler) activityHTML(ctx *gin.Context) {
	activity, err := g.userUC.GetUsersActivity()
	if err != nil {
		panic(err) // TODO: log error
	}

	ctx.JSON(http.StatusOK, activity)
	// return ctx.Render("users_activity", fiber.Map{"usersActivity": activity})
}
