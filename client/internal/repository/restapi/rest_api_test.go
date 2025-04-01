package restapi

import (
	"testing"

	"Filmer/client/config"
	"Filmer/client/internal/repository"
)

var userAuthData = repository.AuthIn{
	Email:    "user1@gmail.com",
	Password: "qwerty123",
}
var authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM0NTI1NzIsInN1YiI6IjViNjk3OGVkLTZjYjQtNDk2Zi04ZGIzLTYzY2RlZDc5YTg1YyJ9.Fk_0FFHJri858Qwg5k4m7H7z8T3lAViZtxjfOss1vw4"
var movieID = "baab6f0f-e954-454f-87b2-ca50f4a4d899" // f86f517b-e509-42d4-a1f3-b8cc769cb938

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

func TestGetSearchMovies(t *testing.T) {
	t.Log("Test search movies with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	// create query-params struct
	query := repository.SearchMoviesIn{
		Keyword: "матрица",
		Page:    1,
	}

	apiResp, err := api.GetSearchMovies(authToken, &query)
	if err != nil {
		t.Fatalf("Searching movies failed: %v", err)
	}
	t.Logf("Successfully search movies: %+v", apiResp)
}

func TestGetMovie(t *testing.T) {
	t.Log("Test get full movie info with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	apiResp, err := api.GetMovie(authToken, 301)
	if err != nil {
		t.Fatalf("Get full movie info failed: %v", err)
	}
	t.Logf("Successfully got full movie info: %+v", apiResp)
}

func TestStarMovie(t *testing.T) {
	t.Log("Test star movie with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	apiResp, err := api.PostStar(authToken, movieID)
	if err != nil {
		t.Fatalf("Star movie failed: %v", err)
	}
	t.Logf("Successfully star movie: %+v", apiResp)
}

func TestGetStared(t *testing.T) {
	t.Log("Test get stared movies with REST API")

	// init api client
	cfg := config.NewConfig()
	api := NewRestAPI(cfg)

	categoryIn := repository.CategoryUserMoviesIn{
		"ratingFrom": []string{"8.6"},
		// "yearFrom":   []string{"1999"},
		// "yearTo":     []string{"1999"},
		// "type":       []string{"фильм"},
		// "genres":     []string{"криминал"},
		// "sortField":  []string{"title"},
		// "sortOrder":  []string{"desc"},
		// "page":       []string{"1"},
	}

	apiResp, err := api.GetStared(authToken, categoryIn)
	if err != nil {
		t.Fatalf("Get stared movies failed: %v", err)
	}
	t.Logf("Successfully got stared movies: %+v", apiResp)
}
