package entity

import (
	"time"

	"github.com/google/uuid"
)


// user movie model
//easyjson:json
// @description user movie received data
type UserMovie struct {
	// user uuid
	UserID		uuid.UUID 	`gorm:"not null;type:uuid;primaryKey" json:"-"`
	// movie uuid
	MovieID		uuid.UUID 	`gorm:"not null;type:uuid;primaryKey" json:"-"`
	// movie list (0 - nothing, 1 - want, 2 - watched)
	Status		int8		`gorm:"not null;type:SMALLINT" json:"status" example:"1"`
	// movie star
	Stared		bool		`gorm:"not null;type:BOOLEAN" json:"stared" example:"true"`
	// last update time
	UpdatedAt	time.Time	`gorm:"not null;type:TIMESTAMP" json:"-"`
	
	User		User	`gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Movie		*Movie	`gorm:"not null;foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie,omitempty"`
}
func (UserMovie) TableName() string {
  return "user_movies"
}


//easyjson:json
type UserMoviesFilter struct {
	RatingFrom	*float64 `json:"ratingFrom,omitempty" query:"ratingFrom" validate:"omitempty,min=0"`
	YearFrom	int `json:"yearFrom,omitempty" query:"yearFrom" validate:"omitempty,min=1500,max=3000"`
	YearTo		int `json:"yearTo,omitempty" query:"yearTo" validate:"omitempty,min=1500,max=3000"`
	Type		string `json:"type,omitempty" query:"type" validate:"omitempty,oneof=фильм сериал видео мини-сериал"`
	Genres		[]string `json:"genres,omitempty" query:"genres" validate:"omitempty"`
}
//easyjson:json
type UserMoviesSort struct {
	SortField 	string `json:"sortField,omitempty" query:"sortField" validate:"omitempty,oneof=title rating year updated_at"`
	SortOrder 	string `json:"sortOrder,omitempty" query:"sortOrder" validate:"omitempty,oneof=asc desc"`
}
//easyjson:json
type UserMoviesPagination struct {
	Page 	int `json:"page" query:"page" validate:"omitempty,min=1"`
	Pages 	int `json:"pages"`
	Total 	int64 `json:"total"`
	Limit 	int `json:"limit"`
}
//easyjson:json
type UserMoviesWithCategory struct {
	UserID		uuid.UUID `json:"-"`
	Category	string `json:"-"`

	// filter fields
	Filter		*UserMoviesFilter `json:"filter"`
	// sort by field
	Sort		*UserMoviesSort `json:"sort"`
	// pagination
	Pagination	*UserMoviesPagination `json:"pagination"`
	// found movies list
	UserMovies	[]UserMovie `json:"films"`
}
