package http

import (
	"fmt"
	"net/http"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/app/constants"
	"Filmer/client/internal/pkg/utils"
	"Filmer/client/internal/repository"
	restAPI "Filmer/client/internal/repository/restapi"
)

// Manager for user subroutes handlers
type userHandlerManager struct {
	cfg *config.Config
	api repository.RestAPI
}

// userHandlerManager constructor
func newUserHandlerManager(cfg *config.Config) *userHandlerManager {
	return &userHandlerManager{
		cfg: cfg,
		api: restAPI.NewRestAPI(cfg),
	}
}

// Render login page
func (hm userHandlerManager) loginGET(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}

// Render sign up page
func (hm userHandlerManager) signUpGET(ctx *fiber.Ctx) error {
	return ctx.Render("signup", fiber.Map{})
}

// Render profile page
func (hm userHandlerManager) profileGET(ctx *fiber.Ctx) error {
	accessToken := ctx.Locals("accessToken").(string)
	// calc hours, minutes and seconds to expiration time
	days, hours, minutes, seconds, err := utils.GetJWTExpirationData(accessToken)
	if err != nil {
		return fmt.Errorf("render profile: get token expiration time: %w", err)
	}
	return ctx.Render("profile", fiber.Map{
		"email":   ctx.Locals("email"),
		"days":    days,
		"hours":   hours,
		"minutes": minutes,
		"seconds": seconds,
	})
}

// Login user via send request to REST API
func (hm userHandlerManager) loginPOST(ctx *fiber.Ctx) error {
	var err error
	var authIn repository.AuthIn

	// parse JSON-body
	if err = ctx.BodyParser(&authIn); err != nil {
		return fmt.Errorf("login: %w", err)
	}
	// send request to REST API
	apiResp, err := hm.api.Login(authIn)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	// get token and email from API response
	accessToken := (*apiResp)["accessToken"].(string)
	email := (*apiResp)["user"].(map[string]any)["email"].(string)
	// set auth and email cookies
	ctx.Cookie(utils.GetAuthCookie(hm.cfg, accessToken))
	ctx.Cookie(utils.GetEmailCookie(hm.cfg, email))

	reidrectURL, err := url.QueryUnescape(ctx.Query(constants.NextQueryParam, "/user/profile"))
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	return ctx.Redirect(reidrectURL, http.StatusSeeOther)
}

// Sign up user via send request to REST API
func (hm userHandlerManager) signUpPOST(ctx *fiber.Ctx) error {
	var err error
	var authIn repository.AuthIn

	// parse JSON-body
	if err = ctx.BodyParser(&authIn); err != nil {
		return fmt.Errorf("sign up: %w", err)
	}
	// check password and password confirm is equal
	if authIn.Password != authIn.PasswordConfirm {
		return fmt.Errorf("sign up: %w", fiber.NewError(http.StatusBadRequest, "passwords do not match"))
	}

	// send request to REST API
	apiResp, err := hm.api.SignUp(authIn)
	if err != nil {
		return fmt.Errorf("sign up: %w", err)
	}
	// get token and email from API response
	accessToken := (*apiResp)["accessToken"].(string)
	email := (*apiResp)["user"].(map[string]any)["email"].(string)
	// set auth and email cookies
	ctx.Cookie(utils.GetAuthCookie(hm.cfg, accessToken))
	ctx.Cookie(utils.GetEmailCookie(hm.cfg, email))

	return ctx.Redirect("/user/profile", http.StatusSeeOther)
}

// Logout user via send request to REST API
func (hm userHandlerManager) logoutPOST(ctx *fiber.Ctx) error {
	accessToken := ctx.Locals("accessToken").(string)
	// send request to REST API
	err := hm.api.Logout(accessToken)
	if err != nil {
		return fmt.Errorf("logout: %w", err)
	}
	// clear auth and email cookies
	ctx.Cookie(utils.ClearAuthCookie(hm.cfg))
	ctx.Cookie(utils.ClearEmailCookie(hm.cfg))

	return ctx.Redirect("/", http.StatusSeeOther)
}
