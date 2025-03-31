package repository

// interface for REST API client
type RestAPI interface {
	// user movie (GET)
	GetMovie(authToken string, kinopoiskID int) (*APIResponse, error)
	GetStared(authToken string, queryParams CategoryUserMoviesIn) (*APIResponse, error)
	GetWant(authToken string, queryParams CategoryUserMoviesIn) (*APIResponse, error)
	GetWatched(authToken string, queryParams CategoryUserMoviesIn) (*APIResponse, error)
	// user movie (POST)
	PostStar(authToken string, movieID string) (*APIResponse, error)
	PostUnstar(authToken string, movieID string) (*APIResponse, error)
	PostClear(authToken string, movieID string) (*APIResponse, error)
	PostWant(authToken string, movieID string) (*APIResponse, error)
	PostWatched(authToken string, movieID string) (*APIResponse, error)
	// movie
	GetSearchMovies(authToken string, queryParams *SearchMoviesIn) (*APIResponse, error)
	// user
	Login(body AuthIn) (*APIResponse, error)
	SignUp(body AuthIn) (*APIResponse, error)
	Logout(authToken string) error
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

// Query-params for get stared || want || watched movies
type CategoryUserMoviesIn map[string][]string

// type CategoryUserMoviesIn struct {
// 	// filter
// 	RatingFrom *string  `query:"ratingFrom,omitempty"`
// 	YearFrom   *string  `query:"yearFrom,omitempty"`
// 	YearTo     *string  `query:"yearTo,omitempty"`
// 	Type       *string  `query:"type,omitempty"`
// 	Genres     []string `query:"genres,omitempty"`
// 	// sort
// 	SortField *string `query:"sortField,omitempty"`
// 	SortOrder *string `query:"sortOrder,omitempty"`
// 	// pagination
// 	Page *string `query:"page,omitempty"`
// }

// struct to parse and return JSON-response from REST API
type APIResponse map[string]any
