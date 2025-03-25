package rest_api

import (
	"testing"

	"Filmer/client/config"
	"Filmer/client/internal/repository"
)

var userAuthData = repository.AuthIn{
	Email:    "user1@gmail.com",
	Password: "qwerty123",
}
var authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDI4MjE1NDQsInN1YiI6IjViNjk3OGVkLTZjYjQtNDk2Zi04ZGIzLTYzY2RlZDc5YTg1YyJ9.wf7S8UvU5mBenzdHIRwqwxjs0zs7ihC2qaUkxwPM1DY"

func TestSignUp(t *testing.T) {
	t.Log("Test sign up with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	apiResp, err := api.SignUp(userAuthData)
	if err != nil {
		t.Fatalf("Sign up failed: %v", err)
	}
	t.Logf("Successfully sign up: %#v", apiResp)
	authToken = (*apiResp)["accessToken"].(string)
}

func TestLogout(t *testing.T) {
	t.Log("Test logout with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	err := api.Logout(authToken)
	if err != nil {
		t.Fatalf("Logout failed: %v", err)
	}
	t.Log("Successfully logout!")
}

func TestLogin(t *testing.T) {
	t.Log("Test login with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	apiResp, err := api.Login(userAuthData)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	t.Logf("Successfully login: %+v", apiResp)
	authToken = (*apiResp)["accessToken"].(string)
}
