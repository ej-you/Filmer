package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/app/constants"
	"Filmer/client/internal/repository"
	restAPI "Filmer/client/internal/repository/restapi"
)

// Manager for staff subroutes handlers.
type staffHandlerManager struct {
	cfg *config.Config
	api repository.RestAPI
}

func newStaffHandlerManager(cfg *config.Config) *staffHandlerManager {
	return &staffHandlerManager{
		cfg: cfg,
		api: restAPI.NewRestAPI(cfg),
	}
}

// Render person page
func (p staffHandlerManager) personGET(ctx *fiber.Ctx) error {
	var err error
	// parse person ID from path
	personID, err := ctx.ParamsInt("personID")
	if err != nil {
		return fmt.Errorf("person: %w", err)
	}
	if personID <= 0 {
		return fiber.NewError(http.StatusBadRequest, "invalid person ID was given")
	}
	// get access token
	accessToken := ctx.Locals(constants.LocalsKeyAccessToken).(string)

	// send request to REST API
	apiResp, err := p.api.GetPerson(accessToken, personID)
	if err != nil {
		return fmt.Errorf("person: %w", err)
	}

	return ctx.Render("person", fiber.Map(
		map[string]any{"person": *apiResp},
	))
}
