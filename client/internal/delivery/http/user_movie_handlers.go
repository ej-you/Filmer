package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/repository"
	restAPI "Filmer/client/internal/repository/rest_api"
)

// Manager for user movie subroutes handlers
type userMovieHandlerManager struct {
	cfg *config.Config
	api repository.RestAPI
}

// userMovieHandlerManager constructor
func newUserMovieHandlerManager(cfg *config.Config) *userMovieHandlerManager {
	return &userMovieHandlerManager{
		cfg: cfg,
		api: restAPI.NewRestAPI(cfg),
	}
}

// Render movie page
func (hm userMovieHandlerManager) movieGET(ctx *fiber.Ctx) error {
	var err error
	// parse movie ID from path
	movieID, err := ctx.ParamsInt("movieID")
	if err != nil {
		return fmt.Errorf("movie: %w", err)
	}
	if movieID <= 0 {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	apiResp, err := hm.api.GetMovie(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("movie: %w", err)
	}

	// transform status field to stared, want and watched
	(*apiResp)["want"] = false
	(*apiResp)["watched"] = false
	userMovieStatus := (*apiResp)["status"].(float64)
	if userMovieStatus == 1 {
		(*apiResp)["want"] = true
	} else if userMovieStatus == 2 {
		(*apiResp)["watched"] = true
	}

	// process last update date
	movieData := (*apiResp)["movie"].(map[string]any)
	updatedAt := movieData["updatedAt"].(string)
	if updatedAt != "" {
		movieData["updatedAt"] = updatedAt[:10]
		(*apiResp)["movie"] = movieData
	}

	return ctx.Render("movie", fiber.Map(*apiResp))
}
