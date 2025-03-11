package entity

import (
	"github.com/google/uuid"
)

// movie genre model
// @description movie genre
//
//easyjson:json
type Genre struct {
	// movie uuid
	MovieID uuid.UUID `gorm:"not null;type:uuid;primaryKey" json:"-"`
	// movie genre
	Genre string `gorm:"not null;type:VARCHAR(50);primaryKey" json:"genre" example:"боевик"`

	Movie Movie `gorm:"not null;foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (Genre) TableName() string {
	return "genres"
}
