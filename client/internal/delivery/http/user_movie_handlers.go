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
	// parse kinopoisk ID from path
	kinopoiskID, err := ctx.ParamsInt("kinopoiskID")
	if err != nil {
		return fmt.Errorf("movie: %w", err)
	}
	if kinopoiskID <= 0 {
		return fiber.NewError(400, "invalid kinopoisk ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	apiResp, err := hm.api.GetMovie(accessToken, kinopoiskID)
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

// Star user movie via send request to REST API
func (hm userMovieHandlerManager) starPOST(ctx *fiber.Ctx) error {
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue("next", "/")
	// parse movie ID from path
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	_, err := hm.api.PostStar(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("star movie: %w", err)
	}
	return ctx.Redirect(next, 303)
}

// Unstar user movie via send request to REST API
func (hm userMovieHandlerManager) unstarPOST(ctx *fiber.Ctx) error {
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue("next", "/")
	// parse movie ID from path
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	_, err := hm.api.PostUnstar(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("unstar movie: %w", err)
	}
	return ctx.Redirect(next, 303)
}

// Clear user movie categories via send request to REST API
func (hm userMovieHandlerManager) clearPOST(ctx *fiber.Ctx) error {
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue("next", "/")
	// parse movie ID from path
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	_, err := hm.api.PostClear(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("clear movie category: %w", err)
	}
	return ctx.Redirect(next, 303)
}

// Set "want" user movie category via send request to REST API
func (hm userMovieHandlerManager) wantPOST(ctx *fiber.Ctx) error {
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue("next", "/")
	// parse movie ID from path
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	_, err := hm.api.PostWant(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("set want movie category: %w", err)
	}
	return ctx.Redirect(next, 303)
}

// Set "watched" user movie via send request to REST API
func (hm userMovieHandlerManager) watchedPOST(ctx *fiber.Ctx) error {
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue("next", "/")
	// parse movie ID from path
	movieID := ctx.Params("movieID")
	if movieID == "" {
		return fiber.NewError(400, "invalid movie ID was given")
	}
	// get access token
	accessToken := ctx.Locals("accessToken").(string)

	// send request to REST API
	_, err := hm.api.PostWatched(accessToken, movieID)
	if err != nil {
		return fmt.Errorf("set watched movie category: %w", err)
	}
	return ctx.Redirect(next, 303)
}
