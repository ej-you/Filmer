// Package entity contains all app entities.
package entity

import (
	"fmt"

	"github.com/google/uuid"
)

// Movie model
type Movie struct {
	// movie ID
	id uuid.UUID
}

// NewMovie parse movie data from given msg and creates new Movie instance.
func NewMovie(rawData []byte) (*Movie, error) {
	var (
		err   error
		movie = &Movie{}
	)
	// parse movie id
	movie.id, err = uuid.FromBytes(rawData)
	if err != nil {
		return nil, fmt.Errorf("parse movie id: %w", err)
	}
	return movie, nil
}

func (m Movie) ID() string {
	return m.id.String()
}
