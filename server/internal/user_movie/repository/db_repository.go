package repository

import (
	"errors"
	"fmt"
	"math"

	"gorm.io/gorm"

	"Filmer/server/internal/entity"
	httpError "Filmer/server/pkg/http_error"
	
	userMovie "Filmer/server/internal/user_movie"
)


const paginationLimit = 10 // user movies per page


// userMovie.Repository interface implementation
type userMovieRepository struct {
	dbClient *gorm.DB
}

// userMovie.Repository constructor
// Returns userMovie.Repository interface
func NewRepository(dbClient *gorm.DB) userMovie.Repository {
	return &userMovieRepository{
		dbClient: dbClient,
	}
}

// Get user movie
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Fill given userMovie struct
// Returns true, if user movie was found in DB, else false
func (this userMovieRepository) GetUserMovie(userMovie *entity.UserMovie) (bool, error) {
	// find user movie in DB
	selectResult := this.dbClient.
		Where("movie_id = ? AND user_id = ?", userMovie.MovieID, userMovie.UserID).
		First(userMovie)
	if err := selectResult.Error; err != nil {
		// if NOT "Not found" error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, httpError.NewHTTPError(500, "failed to get user movie", err)
		}
		return false, nil
	}
	return true, nil
}

// Get user movie or create it if it does not exist
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Fill given userMovie struct
func (this userMovieRepository) FindOrCreateUserMovie(userMovie *entity.UserMovie) error {
	selectResult := this.dbClient.Where(userMovie).FirstOrCreate(userMovie)
	if err := selectResult.Error; err != nil {
		return httpError.NewHTTPError(500, "failed to find or create user movie", err)
	}
	return nil
}


// Set stared field of user movie to newStared value
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Found user movie to update by given userMovie struct
// Fill given userMovie struct
func (this userMovieRepository) UpdateUserMovieStared(userMovie *entity.UserMovie, newStared bool) error {
	// update stared
	updateResult := this.dbClient.Model(userMovie).Update("stared", newStared)
	if err := updateResult.Error; err != nil {
		return httpError.NewHTTPError(500, "failed to update stared field of user movie", err)
	}
	return nil
}

// Set status field of user movie to newStatus value
// Must be presented movie ID (userMovie.MovieID) and user ID (userMovie.UserID)
// Found user movie to update by given userMovie struct
// Fill given userMovie struct
func (this userMovieRepository) UpdateUserMovieStatus(userMovie *entity.UserMovie, newStatus int8) error {
	// update status
	updateResult := this.dbClient.Model(userMovie).Update("status", newStatus)
	if err := updateResult.Error; err != nil {
		return httpError.NewHTTPError(500, "failed to update status field of user movie", err)
	}
	return nil
}


// Get user movies in given category (stared || want || watched)
// Must be presented category (userMoviesWithCategory.Category) and user ID (userMoviesWithCategory.UserID)
// Fill given userMoviesWithCategory struct
func (this userMovieRepository) GetUserMoviesWithCategory(userMoviesWithCategory *entity.UserMoviesWithCategory) error {
	// define select condition for given category
	var categoryCond string
	switch userMoviesWithCategory.Category {
		case "stared":
			categoryCond = "stared = true"
		case "want":
			categoryCond = "status = 1"
		case "watched":
			categoryCond = "status = 2"
	}

	// base select params
	selectQuery := this.dbClient.
		Table("user_movies").
		Distinct("user_movies.*, movies.title, movies.rating, movies.year").
		InnerJoins("INNER JOIN movies ON user_movies.movie_id = movies.id").
		Where(categoryCond).
		Where("user_id = ?", userMoviesWithCategory.UserID)
	selectQuery = this.addSort(selectQuery, userMoviesWithCategory.Sort)
	selectQuery = this.addFilter(selectQuery, userMoviesWithCategory.Filter)
	selectQuery = this.addPagination(selectQuery, userMoviesWithCategory.Pagination)
	// add preloading data from dependent tables
	selectQuery = selectQuery.
		Preload("Movie", func(dbOmit *gorm.DB) *gorm.DB {
		    // omit heavy fields from preloading
		    return dbOmit.Omit("WebURL", "MovieLength", "Description", "Staff")
		}).
		Preload("Movie.Genres")

	// do select query
	selectResult := selectQuery.Debug().Find(&userMoviesWithCategory.UserMovies)
	if err := selectResult.Error; err != nil {
		return httpError.NewHTTPError(500, "failed to get user movies with category", err)
	}
	return nil
}

// Add sort to select query
func (this *userMovieRepository) addSort(selectQuery *gorm.DB, sort *entity.UserMoviesSort) *gorm.DB {
	// sort by updated_at field if sorf field is not defined
	if sort.SortField == "" {
		sort.SortField = "updated_at"
	}

	// set desc order if sorf field is updated_at and sort order is not defined
	if sort.SortField == "updated_at" && sort.SortOrder == "" {
		sort.SortOrder = "desc"
	} else {
		// set asc order if sorf field is defined but sort order is not defined
		if sort.SortOrder == "" {
			sort.SortOrder = "asc"
		}
	}
	return selectQuery.Order(fmt.Sprintf("%s %s", sort.SortField, sort.SortOrder))
}
// Add filter to select query
func (this *userMovieRepository) addFilter(selectQuery *gorm.DB, filter *entity.UserMoviesFilter) *gorm.DB {
	if filter.RatingFrom != nil {
		selectQuery = selectQuery.Where("rating >= ?", filter.RatingFrom)
	}
	if filter.YearFrom != 0 {
		selectQuery = selectQuery.Where("year >= ?", filter.YearFrom)
	}
	if filter.YearTo != 0 {
		selectQuery = selectQuery.Where("year <= ?", filter.YearTo)
	}
	if filter.Type != "" {
		selectQuery = selectQuery.Where("type = ?", filter.Type)
	}
	// have at least one genre from given genres
	if len(filter.Genres) > 0 {
		selectQuery = selectQuery.
			InnerJoins(`INNER JOIN genres ON genres.movie_id = movies.id`).
			Where(`genres.genre IN ?`, filter.Genres)
	}
	return selectQuery
}
// Add pagination to select query (add after all filters)
func (this *userMovieRepository) addPagination(selectQuery *gorm.DB, pagination *entity.UserMoviesPagination) *gorm.DB {
	// set page = 1 if page is not defined
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	// set amount limit for page
	pagination.Limit = paginationLimit
	
	// get all user movies amount (suitable for filters)
	selectQuery.Count(&pagination.Total)
	// calc pages amount
	pagination.Pages = int(math.Ceil(float64(pagination.Total) / float64(pagination.Limit)))

	// add pagination params to select query
	return selectQuery.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)
}
