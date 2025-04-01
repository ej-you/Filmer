package http

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/app/constants"
	"Filmer/client/internal/pkg/utils"
	"Filmer/client/internal/repository"
	restAPI "Filmer/client/internal/repository/restapi"
)

const (
	categoryTemplate = "stared_want_watched" // for stared, want and watched render template
	staredCategory   = "stared"
	wantCategory     = "want"
	watchedCategory  = "watched"
	starCategory     = "star"
	unstarCategory   = "unstar"
	clearCategory    = "clear"
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
		return fiber.NewError(http.StatusBadRequest, "invalid kinopoisk ID was given")
	}
	// get access token
	accessToken := ctx.Locals(constants.LocalsKeyAccessToken).(string)

	// send request to REST API
	apiResp, err := hm.api.GetMovie(accessToken, kinopoiskID)
	if err != nil {
		return fmt.Errorf("movie: %w", err)
	}

	// transform status field to stared, want and watched
	(*apiResp)["want"] = false
	(*apiResp)["watched"] = false
	userMovieStatus := (*apiResp)["status"].(float64)
	if userMovieStatus == constants.UserMovieStatusWant {
		(*apiResp)["want"] = true
	} else if userMovieStatus == constants.UserMovieStatusWatched {
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

// Render stared page
func (hm userMovieHandlerManager) staredGET(ctx *fiber.Ctx) error {
	return hm.categoryGET(ctx, staredCategory)
}

// Render want page
func (hm userMovieHandlerManager) wantGET(ctx *fiber.Ctx) error {
	return hm.categoryGET(ctx, wantCategory)
}

// Render watched page
func (hm userMovieHandlerManager) watchedGET(ctx *fiber.Ctx) error {
	return hm.categoryGET(ctx, watchedCategory)
}

// Star user movie via send request to REST API
func (hm userMovieHandlerManager) starPOST(ctx *fiber.Ctx) error {
	return hm.categoryPOST(ctx, starCategory)
}

// Unstar user movie via send request to REST API
func (hm userMovieHandlerManager) unstarPOST(ctx *fiber.Ctx) error {
	return hm.categoryPOST(ctx, unstarCategory)
}

// Clear user movie categories via send request to REST API
func (hm userMovieHandlerManager) clearPOST(ctx *fiber.Ctx) error {
	return hm.categoryPOST(ctx, clearCategory)
}

// Set "want" user movie category via send request to REST API
func (hm userMovieHandlerManager) wantPOST(ctx *fiber.Ctx) error {
	return hm.categoryPOST(ctx, wantCategory)
}

// Set "watched" user movie via send request to REST API
func (hm userMovieHandlerManager) watchedPOST(ctx *fiber.Ctx) error {
	return hm.categoryPOST(ctx, watchedCategory)
}

func (hm userMovieHandlerManager) categoryGET(ctx *fiber.Ctx, category string) error {
	var err error
	// parse query-params
	categoryIn, err := utils.GetCategoryQueryParams(ctx)
	if err != nil {
		return fmt.Errorf("get %s: %w", category, err)
	}
	// get access token
	accessToken := ctx.Locals(constants.LocalsKeyAccessToken).(string)

	// send request to REST API
	apiResp, err := hm.api.GetCategory(accessToken, category, categoryIn)
	if err != nil {
		return fmt.Errorf("get %s: %w", category, err)
	}
	(*apiResp)["title"] = category
	return ctx.Render(categoryTemplate, fiber.Map(*apiResp))
}

func (hm userMovieHandlerManager) categoryPOST(ctx *fiber.Ctx, category string) error {
	// parse movie ID from path
	movieID, err := utils.GetMovieIDPathParam(ctx)
	if err != nil {
		return fmt.Errorf("set %s user movie category: %w", category, err)
	}
	// get access token
	accessToken := ctx.Locals(constants.LocalsKeyAccessToken).(string)

	// send request to REST API
	if _, err = hm.api.PostCategory(accessToken, category, movieID); err != nil {
		return fmt.Errorf("set %s user movie category: %w", category, err)
	}
	// parse "next" form param to redirect after POST will be processed
	next := ctx.FormValue(constants.NextQueryParam, "/")
	return ctx.Redirect(next, http.StatusSeeOther)
}
