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

// Manager for movie subroutes handlers
type movieHandlerManager struct {
	cfg *config.Config
	api repository.RestAPI
}

// movieHandlerManager constructor
func newMovieHandlerManager(cfg *config.Config) *movieHandlerManager {
	return &movieHandlerManager{
		cfg: cfg,
		api: restAPI.NewRestAPI(cfg),
	}
}

// Render search movies page
func (hm movieHandlerManager) searchGET(ctx *fiber.Ctx) error {
	var err error
	searchMoviesIn := new(repository.SearchMoviesIn)

	accessToken := ctx.Locals(constants.LocalsKeyAccessToken).(string)
	// parse query-params
	if err = ctx.QueryParser(searchMoviesIn); err != nil {
		return fmt.Errorf("search movies: %w", err)
	}
	// if need load page without searching movies
	if searchMoviesIn.Keyword == "" {
		return ctx.Render("search", fiber.Map{})
	}
	// if page is not specified (or not valid)
	if searchMoviesIn.Page <= 0 {
		searchMoviesIn.Page = 1
	}

	// send request to REST API
	apiResp, err := hm.api.GetSearchMovies(accessToken, searchMoviesIn)
	if err != nil {
		return fmt.Errorf("search movies: %w", err)
	}

	// if requested page value more than existing pages amount
	reqPage := (*apiResp)["page"].(float64)
	pages := (*apiResp)["pages"].(float64)
	if reqPage > pages {
		return fiber.NewError(http.StatusNotFound, "search movie page with given page param not fount")
	}
	return ctx.Render("search", fiber.Map(*apiResp))
}
