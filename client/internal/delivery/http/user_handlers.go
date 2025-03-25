package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"Filmer/client/config"
	"Filmer/client/internal/pkg/utils"
	"Filmer/client/internal/repository"
	restAPI "Filmer/client/internal/repository/rest_api"
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
	d, h, m, s, err := utils.GetJWTExpirationData(accessToken)
	if err != nil {
		return fmt.Errorf("render profile: get token expiration time", err)
	}
	return ctx.Render("profile", fiber.Map{
		"email":   ctx.Locals("email"),
		"days":    d,
		"hours":   h,
		"minutes": m,
		"seconds": s,
	})
}

// Login user via send request to REST API
func (hm userHandlerManager) loginPOST(ctx *fiber.Ctx) error {
	var err error
	var authIn repository.AuthIn

	// parse JSON-body
	if err = ctx.BodyParser(&authIn); err != nil {
		return fmt.Errorf("login", err)
	}
	// send request to REST API
	apiResp, err := hm.api.Login(authIn)
	if err != nil {
		return fmt.Errorf("login", err)
	}
	// get token and email from API response
	accessToken := (*apiResp)["accessToken"].(string)
	email := (*apiResp)["user"].(map[string]any)["email"].(string)
	// set auth and email cookies
	ctx.Cookie(utils.GetAuthCookie(hm.cfg, accessToken))
	ctx.Cookie(utils.GetEmailCookie(hm.cfg, email))

	return ctx.Redirect("/user/profile", 303)
}

// Sign up user via send request to REST API
func (hm userHandlerManager) signUpPOST(ctx *fiber.Ctx) error {
	var err error
	var authIn repository.AuthIn

	// parse JSON-body
	if err = ctx.BodyParser(&authIn); err != nil {
		return fmt.Errorf("sign up", err)
	}
	// check password and password confirm is equal
	if authIn.Password != authIn.PasswordConfirm {
		return fmt.Errorf("sign up: %w", fiber.NewError(400, "passwords do not match"))
	}

	// send request to REST API
	apiResp, err := hm.api.SignUp(authIn)
	if err != nil {
		return fmt.Errorf("sign up", err)
	}
	// get token and email from API response
	accessToken := (*apiResp)["accessToken"].(string)
	email := (*apiResp)["user"].(map[string]any)["email"].(string)
	// set auth and email cookies
	ctx.Cookie(utils.GetAuthCookie(hm.cfg, accessToken))
	ctx.Cookie(utils.GetEmailCookie(hm.cfg, email))

	return ctx.Redirect("/user/profile", 303)
}

// Logout user via send request to REST API
func (hm userHandlerManager) logoutPOST(ctx *fiber.Ctx) error {
	accessToken := ctx.Locals("accessToken").(string)
	// send request to REST API
	err := hm.api.Logout(accessToken)
	if err != nil {
		return fmt.Errorf("logout", err)
	}
	// clear auth and email cookies
	ctx.Cookie(utils.ClearAuthCookie(hm.cfg))
	ctx.Cookie(utils.ClearEmailCookie(hm.cfg))

	return ctx.Redirect("/", 303)
}
