package http

// @description data for change user password
//
//easyjson:json
type changePasswordIn struct {
	// current password
	CurrentPassword string `json:"currentPassword" validate:"required,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
	// new password
	NewPassword string `json:"newPassword" validate:"required,min=8,max=40" example:"123qwerty" minLength:"8" maxLength:"40"`
}
