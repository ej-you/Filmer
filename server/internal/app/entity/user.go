package entity

import (
	"time"

	"github.com/google/uuid"
)

// @description user model
//
//easyjson:json
type User struct {
	// user uuid
	ID uuid.UUID `gorm:"not null;type:uuid;default:uuid_generate_v4();primaryKey" json:"-"`
	// user email
	Email string `gorm:"not null;type:VARCHAR(100);uniqueIndex:uni_users_email" json:"email" example:"user@gmail.com"`
	// user password hash
	Password []byte `gorm:"not null;type:BYTES" json:"-"`
	// create account date
	CreatedAt time.Time `gorm:"not null;type:TIMESTAMP" json:"-"`

	UserMovies []UserMovie `gorm:"foreignKey:UserID" json:"-"`
}

func (User) TableName() string {
	return "users"
}

// @description received data about user with token
//
//easyjson:json
type UserWithToken struct {
	// user data
	User *User `json:"user"`
	// access token, generated after success auth
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA3NjA0MzIsInVzZXJJRCI6IjU4MzU0ZGJhLWUyY2MtNDI4My04MjAxLWNjZGRiZTkzOGExNSJ9.VwA9d9lao0Xgl5c3H9VNM8JVtUKDs47YEItb6MZlkCw"`
}

// @description received data about one user activity (for admin-panel)
//
//easyjson:json
type UserActivity struct {
	// user email
	Email string `json:"email" example:"user@gmail.com"`
	// create user datetime (RFC3339 format)
	CreatedAt time.Time `json:"createdAt" example:"2025-02-02T22:00:07.225526936Z"`

	// stared user movies
	Stared int `json:"stared" example:"5"`
	// user movies in "want" list
	Want int `json:"want" example:"25"`
	// user movies in "watched" list
	Watched int `json:"watched" example:"50"`
}

// @description received data about all users activity
//
//easyjson:json
type UsersActivity []UserActivity
