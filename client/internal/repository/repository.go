package repository

// interface for REST API client
type RestAPI interface {
	// user movie (GET)
	GetMovie(authToken string, kinopoiskID int) (*APIResponse, error)
	GetCategory(authToken, category string, queryParams CategoryUserMoviesIn) (*APIResponse, error)
	PostCategory(authToken, category, movieID string) (*APIResponse, error)
	// movie
	GetSearchMovies(authToken string, queryParams *SearchMoviesIn) (*APIResponse, error)
	// user
	Login(body AuthIn) (*APIResponse, error)
	SignUp(body AuthIn) (*APIResponse, error)
	Logout(authToken string) error
	ChangePassword(authToken string, body ChangePasswordIn) error
	// staff
	GetPerson(authToken string, kinopoiskID int) (*APIResponse, error)
}

// JSON-body for login && signup
type AuthIn struct {
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password-confirm,omitempty"` // for sign up only
}

// Query-params for search movies
type SearchMoviesIn struct {
	Keyword string `query:"q"`
	Page    int    `query:"page"`
}

// JSON-body for change user password
type ChangePasswordIn struct {
	CurrentPassword string `form:"current-password" json:"currentPassword"`
	NewPassword     string `form:"new-password" json:"newPassword"`
}

// Query-params for get stared || want || watched movies
type CategoryUserMoviesIn map[string][]string

// struct to parse and return JSON-response from REST API
type APIResponse map[string]any
