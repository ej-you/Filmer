package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	httpError "Filmer/server/internal/pkg/http_error"
)

// Encode given password
// Returns encoded password like a hash
func EncodePassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		// probably, password is too long
		return nil, httpError.NewHTTPError(http.StatusBadRequest, "failed to encode password", err)
	}
	return hash, nil
}

// Check the given password is equal to its hash from DB
// Returns true, if password is equal
func PasswordIsCorrect(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
