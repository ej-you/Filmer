// Package password provides functions for working passwords and its hashes.
package password

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"Filmer/server/internal/pkg/httperror"
)

// Encode encodes given password
// Returns encoded password like a hash
func Encode(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		// probably, password is too long
		return nil, httperror.New(http.StatusBadRequest, "failed to encode password", err)
	}
	return hash, nil
}

// IsCorrect checks the given password is equal to its hash from DB
// Returns true, if password is equal
func IsCorrect(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
