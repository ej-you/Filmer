package http

// @description data for sign up OR login user
//
//easyjson:json
type authIn struct {
	// user email
	Email string `json:"email" validate:"required,max=100" example:"user@gmail.com" maxLength:"100"`
	// user password
	Password string `json:"password" validate:"required,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
}
