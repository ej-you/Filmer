// Package entity contains all app entities.
package entity

import "time"

const _datetimeFormat = "02 Jan 2006 15:04:05" // for PrettyCreatedAt field

type UserActivity struct {
	// user email
	Email string `json:"email"`
	// create user datetime (RFC3339 format)
	CreatedAt time.Time `json:"createdAt"`
	// formatted datetime for output
	PrettyCreatedAt string `json:"prettyCreatedAt"`

	// stared user movies
	Stared int `json:"stared"`
	// user movies in "want" list
	Want int `json:"want"`
	// user movies in "watched" list
	Watched int `json:"watched"`
}

// FormatCreatedAt formats CreatedAt user field according to
// _datetimeFormat and saves result to PrettyCreatedAt field.
func (u *UserActivity) FormatCreatedAt() {
	u.PrettyCreatedAt = u.CreatedAt.Format(_datetimeFormat)
}

type UsersActivity []UserActivity
